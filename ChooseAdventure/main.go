package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	template := template.Must(template.ParseFiles("layout/layout.html"))
	stories := getStories("stories.json")
	sh := &StoryHandler{Stories: stories, StoryTemplate: template}

	http.Handle("/", sh)
	http.Handle("/layout/", http.StripPrefix("/layout/", http.FileServer(http.Dir("layout"))))

	log.Println("Listening...")
	http.ListenAndServe(":8090", nil)

}

func getStories(filename string) map[string]Arc {
	jsonFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()
	var stories map[string]Arc
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &stories)

	return stories
}

func (sh *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := strings.Replace(r.URL.Path, "/", "", 1)
	_, contains := sh.Stories[requestPath]
	data := sh.Stories["intro"]

	if contains {
		data = sh.Stories[requestPath]
	}

	sh.StoryTemplate.Execute(w, data)
}

type StoryHandler struct {
	Stories       map[string]Arc
	StoryTemplate *template.Template
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
