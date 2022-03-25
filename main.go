package main

import (
	"fmt"
	"net/http"
	upload_api "read-pdf/src/upload"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/upload", upload_api.UploadFile).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
