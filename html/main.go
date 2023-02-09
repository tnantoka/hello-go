package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func printTagNames(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Println(node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		printTagNames(child)
	}
}

func main() {
	res, _ := http.Get("https://tnantoka.com/")
	defer res.Body.Close()

	root, _ := html.Parse(res.Body)
	printTagNames(root)
}
