package main

import (
	"fmt"
	"html/template"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/chai2010/webp"
)

const uploadDir = "./uploads"

type PageData struct {
	Images      []UploadedImage
	Message     string
	MessageType string
}

type UploadedImage struct {
	Name string
	URL  string
}

func main() {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create upload directory: %v", err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))

	fmt.Println("Server running: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	images, err := listUploadedImages()
	if err != nil {
		log.Printf("failed to list uploaded images: %v", err)
		http.Error(w, "failed to load gallery", http.StatusInternalServerError)
		return
	}

	data := PageData{Images: images}
	switch r.URL.Query().Get("status") {
	case "success":
		data.Message = "Imagem enviada com sucesso."
		data.MessageType = "success"
	case "error":
		data.Message = "Falha ao enviar imagem. Veja o erro no terminal."
		data.MessageType = "error"
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("failed to parse template: %v", err)
		http.Error(w, "failed to load page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("failed to render template: %v", err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Limita o tamanho do upload para 20MB
	if err := r.ParseMultipartForm(20 << 20); err != nil {
		log.Printf("failed to parse upload form: %v", err)
		redirectUploadError(w, r)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Printf("failed to read uploaded image: %v", err)
		redirectUploadError(w, r)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("failed to decode uploaded image %q: %v", handler.Filename, err)
		redirectUploadError(w, r)
		return
	}

	name := strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename))
	name = fmt.Sprintf("%s_%d.webp", sanitize(name), time.Now().Unix())

	outPath := filepath.Join(uploadDir, name)

	out, err := os.Create(outPath)
	if err != nil {
		log.Printf("failed to create uploaded image %q: %v", outPath, err)
		redirectUploadError(w, r)
		return
	}
	defer out.Close()

	// Converte a imagem para WebP com qualidade de 80
	if err := webp.Encode(out, img, &webp.Options{Quality: 80}); err != nil {
		log.Printf("failed to encode uploaded image %q as webp: %v", handler.Filename, err)
		redirectUploadError(w, r)
		return
	}

	http.Redirect(w, r, "/?status=success", http.StatusSeeOther)
}

func redirectUploadError(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/?status=error", http.StatusSeeOther)
}

func listUploadedImages() ([]UploadedImage, error) {
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		return nil, err
	}

	images := make([]UploadedImage, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !isImageFile(entry.Name()) {
			continue
		}

		images = append(images, UploadedImage{
			Name: entry.Name(),
			URL:  "/uploads/" + entry.Name(),
		})
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].Name > images[j].Name
	})

	return images, nil
}

func isImageFile(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".webp", ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

func sanitize(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "_")
	return s
}
