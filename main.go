package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (page * Page) save() error {
	filename := page.Title + ".txt" //Remember to refer to page.Title 
	return os.WriteFile(filename, page.Body, 0600)
}

func load_page(title string) (*Page, error){
	filename := title + ".txt"
	body, error := os.ReadFile(filename)

	if error != nil {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}

func view_handler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	page, _ := load_page(title)
	fmt.Fprintf(writer, "<h1>%s</h1> <div>%s</div>", page.Title, page.Body) // It's FPrintf
}

func main() {
	http.HandleFunc("/view/", view_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
