package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	r := getHTML("sample.html")
	z := html.NewTokenizer(r)

	for {

		tokenType := z.Next()

		if tokenType == html.ErrorToken {
			err := z.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", z.Err())
		}

		if isStartAnchor(z, tokenType) {
			href := getAnchorLink(z)

			fmt.Println(href)

		}
	}
}

func getAnchorLink(n *html.Tokenizer) string {
	return getAttributeValue(n, "href")
}

func getAttributeValue(n *html.Tokenizer, attributeName string) string {

	for {
		attr, attrVal, hasMoreAttr := n.TagAttr()

		if attr != nil && attrVal != nil {
			if strings.ToLower(string(attr)) == attributeName {
				return string(attrVal)
			}
		}

		if !hasMoreAttr {
			break
		}
	}
	return ""
}

func isStartAnchor(n *html.Tokenizer, tokenType html.TokenType) bool {
	name, _ := n.TagName()
	return string(name) == "a" && tokenType.String() == "StartTag"
}

func getHTML(filePath string) io.Reader {
	data, err := ioutil.ReadFile(filePath)
	check(err)
	reader := strings.NewReader(string(data))

	return reader
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Link struct {
	Href string
	Text string
}
