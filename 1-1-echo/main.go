package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

const addr = ":18888"

func main() {
	http.HandleFunc("/", handler)
	httpServer := http.Server{Addr: addr}
	log.Printf("Start http listening at %s\n", addr)
	log.Println(httpServer.ListenAndServe())
}
