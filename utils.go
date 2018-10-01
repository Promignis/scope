package main

import (
	"fmt"
	"log"
	"net/http"
)

func genError(error string) string {
	return fmt.Sprintf("{error: true, message: '%s'}", error)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func headerToString(header http.Header) string {
	output := ""
	for k, v := range header {
		output += fmt.Sprintf("%s: %s\n", k, v)
	}
	return output
}
