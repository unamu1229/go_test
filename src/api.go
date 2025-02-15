package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		index,
	)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(response http.ResponseWriter, req *http.Request) {
	response.Write([]byte(req.URL.Path))
}
