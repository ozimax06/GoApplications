package main

import (
	"errors"
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
	openAnchorTags := 0
	//innerText := ""
	outerAnchorLink := ""
	//var linkTexts []Link

	for {

		tokenType := z.Next()

		if tokenType == html.ErrorToken {
			err := z.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", z.Err())
		}

		if isAnchor(z) {
			if isStartTag(tokenType) {
				openAnchorTags++

				if openAnchorTags == 1 {
					outerAnchorLink = getAnchorLink(z)
					fmt.Println(outerAnchorLink)
				}
			} else if isEndTag(tokenType) {
				openAnchorTags--

				if openAnchorTags == 0 {

				}
			}
		}

	}
}

func getAnchorLink(n *html.Tokenizer) string {
	result, err := getAttributeValue(n, "href")

	if err != nil {
		result = ""
	}
	return result
}

func getAttributeValue(n *html.Tokenizer, attributeName string) (string, error) {

	for {
		attr, attrVal, hasMoreAttr := n.TagAttr()

		if attr != nil && attrVal != nil {
			if strings.ToLower(string(attr)) == attributeName {
				return string(attrVal), nil
			}
		}

		if !hasMoreAttr {
			break
		}
	}
	return "", errors.New("Attribute couldn't be found")
}

func isAnchor(n *html.Tokenizer) bool {
	name, _ := n.TagName()
	return string(name) == "a"
}

func isText(tokenType html.TokenType) bool {
	return tokenType.String() == "Text"
}

func isStartTag(tokenType html.TokenType) bool {
	return tokenType.String() == "StartTag"
}

func isEndTag(tokenType html.TokenType) bool {
	return tokenType.String() == "EndTag"
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
