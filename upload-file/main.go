package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

	headerType := handler.Header["Content-Type"][0]
	headerTypesArray := []string{"image/png", "image/jpeg", "image/jpg"}
	headerTypes := map[string]string{}
	for _, header := range headerTypesArray {
		headerTypes[header] = header
	}

	// Check the type header of the file
	if headerType != headerTypes[headerType] {
		JSON(w, http.StatusBadRequest, map[string]interface{}{"Message": "The file must be png, jpeg, or jpg"})
		return
	}

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

func deleteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filename := mux.Vars(r)["filename"]

	err := os.Remove("./images/" + filename)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusOK, map[string]interface{}{"Message": "File successfully deleted"})
}

func HandleRoutes() {
	app := mux.NewRouter().StrictSlash(true)

	fmt.Println("Server started at port 8000...")
	app.HandleFunc("/api/upload", uploadFile).Methods("POST")
	app.HandleFunc("/api/delete/{filename}", deleteFile).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", app))
}

func main() {
	HandleRoutes()
}
