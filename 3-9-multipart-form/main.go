package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("name", "John")

	fileName := os.Args[1]
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "text/plain")
	part.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	fw, err := w.CreatePart(part)
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
