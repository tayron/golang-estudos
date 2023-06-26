package method

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {

	var payload bytes.Buffer
	mw := multipart.NewWriter(&payload)

	file, errorOpenFile := os.Open("../testdata/bobs-burgers.jpg")
	if errorOpenFile != nil {
		t.Error(errorOpenFile.Error())
	}

	defer file.Close()

	w, errorCreateFormFile := mw.CreateFormFile("file", "bobs-burgers.jpg")
	if errorCreateFormFile != nil {
		t.Error(errorCreateFormFile.Error())
	}

	_, errorCopyFile := io.Copy(w, file)
	if errorCopyFile != nil {
		t.Error(errorCopyFile.Error())
	}

	mw.Close()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello-you", &payload)
	req.Header.Add("Content-Type", mw.FormDataContentType())
	Upload(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("Status code esperado: %d, status recuperado: %d", http.StatusOK, rr.Code)
	}

	response, errorReadAll := io.ReadAll(rr.Body)
	if errorReadAll != nil {
		t.Error(errorReadAll.Error())
	}

	expect := "Upload feito com sucesso\n"

	if string(response) != expect {
		t.Errorf("Mensagem esperada: %s, status recuperado: %s", expect, string(response))
	}
}
