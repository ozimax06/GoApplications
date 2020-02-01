package main

import (
	"fmt"
	"log"
	"net/http"

	urlshort "github.com/ozan/urlshortener/urlpackages"
)

func main() {
	fallbackHandler := urlshort.DefaultHandler("https://ozanonder.tech")
	yamlFile := urlshort.ReadYAMLFileAsByte("urlpath.yaml")
	handler, err := urlshort.YAMLHandler(yamlFile, fallbackHandler)

	if err != nil {
		fmt.Println("error happened")
	} else {
		http.Handle("/", handler)
		log.Println("Listening...")
		http.ListenAndServe(":8090", nil)
	}

}
