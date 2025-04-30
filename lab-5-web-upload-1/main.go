package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("template/*"))
}

func main() {

	http.HandleFunc("/", handelIndex)
	fmt.Println("Serever start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handelIndex(rsp http.ResponseWriter, req *http.Request) {

	fmt.Println("Request method :", req.Method)

	data := struct {
		File string
	}{
		File: "",
	}

	if req.Method == http.MethodPost {
		f, h, e := req.FormFile("cv")
		if e != nil {
			http.Error(rsp, e.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		b, e := io.ReadAll(f)
		if e != nil {
			http.Error(rsp, e.Error(), http.StatusInternalServerError)
			return
		}
		go writeToDisk(b,h.Filename)
		data.File = string(b)

	}

	fmt.Println(data.File)

	tmp.ExecuteTemplate(rsp, "index.gohtml", data)
}

func writeToDisk(b []byte, name string) {
	f, e := os.Create(filepath.Join("./data/", name))
	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()
	_, err := f.Write(b)

	if err != nil {
		log.Fatal(e)
	}
}
