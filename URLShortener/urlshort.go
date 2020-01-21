package main

import (
	"log"
	"net/http"
)

func main() {
	var urlMaps = map[string]string{
		"/dog":  "https://www.google.com",
		"/cat":  "https://www.cat.com",
		"/ozan": "https://www.yahoo.com",
	}

	defaultHandler := defaultHandle("https://ozanonder.tech")
	handler := MapHandler(urlMaps, defaultHandler)
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

func defaultHandle(defaultURL string) http.Handler {
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
//func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
// TODO: Implement this...
//return nil, nil
//}
