package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("name", "John")

	fileName := os.Args[1]
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(fw, file)
	w.Close()

	resp, err := http.Post("http://localhost:18888", w.FormDataContentType(), &buf)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
