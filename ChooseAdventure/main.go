package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	stories := getStories("stories.json")

	template := template.Must(template.ParseFiles("layout/layout.html"))
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
	var data Arc

	requestPath := strings.Replace(r.URL.Path, "/", "", 1)
	_, contains := sh.Stories[requestPath]

	if contains {
		data = sh.Stories[requestPath]
	} else {
		firstStory, err := getFirstStory(sh.Stories)

		if err != nil {
			panic(err)
		}
		data = sh.Stories[firstStory]
	}

	sh.StoryTemplate.Execute(w, data)
}

func getFirstStory(stories map[string]Arc) (string, error) {

	for storyName := range stories {
		if !isStoryReferenced(storyName, stories) {
			return storyName, nil
		}
	}

	return "", errors.New("first story couldn't be found!")
}

func isStoryReferenced(storyName string, stories map[string]Arc) bool {

	for name, arc := range stories {

		if name == storyName {
			continue
		}

		storyOptions := arc.Options
		for _, option := range storyOptions {

			if option.Arc == storyName {
				return true
			}
		}
	}
	return false
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
