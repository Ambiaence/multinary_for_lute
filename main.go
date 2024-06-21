package main

import (
	"fmt"
	"log"
	"net/http"
)

var elements []string


func view_handler(writer http.ResponseWriter, request *http.Request) {
	word := request.URL.Path[len("/word/"):]
	fmt.Fprint(writer, "<h1>how about now </h1><body>\n") 
	for _, element := range elements {
		fmt.Fprintf(writer, element, word)
	}
}

func construct_embedded_element(website string) string {
	element := ""
	element +="<embed src = \""
	element += website
	element += "%s\" "
	element += "width=\"1000\" height = \"400\" ></embed>\n"
	return element
}

func add_website_to_elements(elements []string, website string) []string {
	element := construct_embedded_element(website)
	elements = append(elements, element)
	return elements 
}

func main() {
	elements = add_website_to_elements(elements, "https://context.reverso.net/translation/german-english/")
	elements = add_website_to_elements(elements, "https://www.dict.cc/?s=")

	http.HandleFunc("/view/", view_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
