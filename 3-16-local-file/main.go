package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	tp := &http.Transport{}
	tp.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	client := http.Client{
		Transport: tp,
	}

	fileName := os.Args[1]
	resp, err := client.Get("file://./" + fileName)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}
