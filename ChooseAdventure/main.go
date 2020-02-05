package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Arc struct {
	Title string
	Story []string
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("rew")
		data := Arc{
			Title: "Title",
			Story: []string{"Story1", "Story2", "Story3"},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8090", nil)
}
