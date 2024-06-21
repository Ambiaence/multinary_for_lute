package main

import (
	"fmt"
	"log"
	"net/http"
)

var elements []string


func view_handler(writer http.ResponseWriter, request *http.Request) {
	word := request.URL.Path[len("/word/"):]
	fmt.Fprint(writer, "<h1> Multinary For Lute </h1><body>\n") 
	for _, element := range elements {
		fmt.Fprintf(writer, element, word)
	}
}

func construct_embedded_element(website string) string {
	element := ""
	element +="<embed src = \""
	element += website
	element += "%s\" "
	element += "width=\"900\" height = \"300\" ></embed>\n"
	return element
}

func add_website_to_elements(elements []string, website string) []string {
	element := construct_embedded_element(website)
	elements = append(elements, element)
	return elements 
}

func main() {
	elements = add_website_to_elements(elements, "https://www.dwds.de/wb/")
	elements = add_website_to_elements(elements, "https://www.dict.cc/?s=")
	elements = add_website_to_elements(elements, "https://context.reverso.net/translation/german-english/")

	http.HandleFunc("/word/", view_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
