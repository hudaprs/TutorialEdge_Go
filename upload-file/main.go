package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Uploading file...")

	// 1. Parse input to type multipart/form-data
	// And set the maximum file size
	r.ParseMultipartForm(10 << 20)

	// 2. Retreive file from posted form-data
	file, handler, err := r.FormFile("file")
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	if file == nil {
		JSON(w, http.StatusBadRequest, map[string]interface{}{"Message": "File is required"})
		return
	}

	// Get the object in handler
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("Mime Header: %+v\n", handler.Header)

	// 3. Write temporary file on our server
	tempFile, err := ioutil.TempFile("images", "upload-*.png")
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	tempFile.Write(fileBytes)

	// 4. Return wheter or not this has been successfull
	JSON(w, http.StatusOK, map[string]interface{}{"Message": "Success upload file"})
}

func HandleRoutes() {
	fmt.Println("Server started at port 8000...")
	http.HandleFunc("/api/upload", uploadFile)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	HandleRoutes()
}
