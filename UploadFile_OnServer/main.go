package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/upload", UploadFile)

	http.ListenAndServe(":8080", nil)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to the Home Page")
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to the Upload Page")

	fmt.Println("r.method = ", r.Method)
	// if method is GET then load form, if not then upload successful message

	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "fileUpload.html", nil)
		return
	} else {
		r.ParseMultipartForm(10)
	}

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	//The defer file.Close() statement ensures that the file is closed after the surrounding function returns, which is a good practice to avoid resource leaks.
	defer file.Close()

	fmt.Printf("fileHeader.Filename = %v\n", fileHeader.Filename)
	fmt.Printf("fileHeader.Size = %v\n", fileHeader.Size)
	fmt.Printf("fileHeader.Header = %v\n", fileHeader.Header)

	// tempFile, err := ioutil.TempFile("images", "upload-*.png")
	contentType := fileHeader.Header["Content-Type"][0]
	fmt.Println("Content Type:", contentType)
	var osFile *os.File
	// func TempFile(dir, pattern string) (f *os.File, err error)
	if contentType == "image/jpeg" {
		osFile, err = ioutil.TempFile("images", "*.jpg")
	} else if contentType == "application/pdf" {
		osFile, err = ioutil.TempFile("PDFs", "*.pdf")
	} else if contentType == "text/javascript" {
		osFile, err = ioutil.TempFile("js", "*.js")
	}
	fmt.Println("error:", err)
	defer osFile.Close()

	// func ReadAll(r io.Reader) ([]byte, error)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// func (f *File) Write(b []byte) (n int, err error)

	osFile.Write(fileBytes)

	fmt.Fprintf(w, "Your File was Successfully Uploaded!\n")

}
