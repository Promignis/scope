package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("u")
	if r.Method == "GET" {

		req, err := http.NewRequest(r.Method, url, nil)
		req.Header = r.Header

		var toRead io.ReadCloser
		resp, err := client.Do(req)
		checkErr(err)

		if resp.Header.Get("Content-Encoding") == "gzip" {
			gzipReader, err := gzip.NewReader(resp.Body)
			checkErr(err)
			toRead = gzipReader
		} else {
			toRead = resp.Body
		}

		defer toRead.Close()

		byteBody, err := ioutil.ReadAll(toRead)
		checkErr(err)
		// TODO: if Accept-Encoding gzip
		// write gzipped data

		fmt.Fprintf(w, string(byteBody))
	} else {
		// write this better
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, genError("POST not yet supported"))
	}
}

func main() {
	http.HandleFunc("/", proxyHandler)

	var port string

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}

	startServer(port)
}

func startServer(port string) {
	fmt.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
