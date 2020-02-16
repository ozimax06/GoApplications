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

		var tagName, tagAttribute, attributeValue []byte
		var hasAttribute, hasMoreAttribute bool
		tagName, hasAttribute = z.TagName()

		fmt.Println(string(tokenType.String()), ": ", string(tagName))

		if hasAttribute {
			for {
				tagAttribute, attributeValue, hasMoreAttribute = z.TagAttr()
				fmt.Println("Attr: ", string(tagAttribute), "AttrVal: ", string(attributeValue))

				if !hasMoreAttribute {
					break
				}
			}
		}

	}
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
