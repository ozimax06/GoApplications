package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func main() {

	ActivateYAMLHandler()

}

//ActivateYAMLHandler tests YANLhandler
func ActivateYAMLHandler() {
	fallbackHandler := DefaultHandler("https://ozanonder.tech")
	yamlFile := ReadYAMLFileAsByte("urlpath.yaml")
	handler, err := YAMLHandler(yamlFile, fallbackHandler)

	if err != nil {
		fmt.Println("error happened")
	} else {
		http.Handle("/", handler)
		log.Println("Listening...")
		http.ListenAndServe(":8090", nil)
	}

}

//ActivateManualMapHandler tests MapHandler
func ActivateManualMapHandler() {

	var urlMaps = map[string]string{
		"/dog":  "https://www.google.com",
		"/cat":  "https://www.cat.com",
		"/ozan": "https://www.yahoo.com",
	}

	fallbackHandler := DefaultHandler("https://ozanonder.tech")
	handler := MapHandler(urlMaps, fallbackHandler)
	http.Handle("/", handler)

	log.Println("Listening...")
	http.ListenAndServe(":8090", nil)
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestPath := r.URL.Path
		_, contains := pathsToUrls[requestPath]

		if contains {
			http.Redirect(w, r, pathsToUrls[requestPath], 301)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}

//DefaultHandler redirects the incoming request to defaultURL
func DefaultHandler(defaultURL string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, defaultURL, 301)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var yamlError error

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var list URLPathList
		requestPath := r.URL.Path

		err := yaml.Unmarshal(yml, &list)
		if err != nil {
			yamlError = errors.New("Error: Parsing YAML file")
			fallback.ServeHTTP(w, r)

		} else {

			yamlError = nil
			for _, mapItem := range list.Maplist {
				if mapItem.Path == requestPath {
					http.Redirect(w, r, mapItem.URL, 301)
				}
			}
			fallback.ServeHTTP(w, r)
		}

	}), yamlError
}

//ReadYAMLFileAsByte reads external yaml file
func ReadYAMLFileAsByte(filename string) []byte {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return data
}

//URLPath represents each item
type URLPath struct {
	Path string
	URL  string
}

//URLPathList is collection of url paths
type URLPathList struct {
	Maplist []URLPath
}
