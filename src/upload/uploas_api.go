package upload_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(response http.ResponseWriter, request *http.Request) {
	// Limit 10MB
	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("myfile")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("File Type: ", handler.Header.Get("Content-Type"))

	// Upload File

	tempFile, err2 := ioutil.TempFile("uploads", "upload-*.pdf")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer tempFile.Close()
	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println(err3)

	}
	tempFile.Write(fileBytes)
	fmt.Println("file upload done")
}
