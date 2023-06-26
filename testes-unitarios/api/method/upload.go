package method

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(rw http.ResponseWriter, req *http.Request) {

	req.ParseMultipartForm(32 << 20)

	file, _, errFormFile := req.FormFile("file")
	if errFormFile != nil {
		http.Error(rw, errFormFile.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	tempFile, errCreateTemp := os.CreateTemp("temp-images", "upload-*.png")
	if errCreateTemp != nil {
		fmt.Println(errCreateTemp.Error())
	}

	defer tempFile.Close()

	fileBytes, errReadFileBytes := io.ReadAll(file)
	if errReadFileBytes != nil {
		fmt.Println(errReadFileBytes.Error())
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(rw, "Upload feito com sucesso\n")
}
