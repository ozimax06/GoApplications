package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//var stories map[string]Arc
	//fmt.Print(stories)

	jsonFile, err := os.Open("stories.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var stories map[string]Arc
	json.Unmarshal([]byte(byteValue), &stories)

	fmt.Println(stories["new-york"].Story[3])

}

func readFromTemplate() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("rew")
		data := "dddo"
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8090", nil)
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
