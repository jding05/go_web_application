package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
)

// describes how page data will be stored in memory
type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename :=title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl, ".html")
	t.Execute(w, p)
}

// The function editHandler loads the page 
// (or, if it doesn't exist, create an empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "view", p)
}

// The function template.ParseFiles will read the contents of edit.html and
//  return a *template.Template.

// The method t.Execute executes the template, writing the generated HTML 
// to the http.ResponseWriter. The .Title and .Body dotted identifiers 
// refer to p.Title and p.Body.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view"):]
	p, _ := loadPage(title)
	renderTemplate(w, "edit", p)
}


// testing
func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http://localhost:8080/view/testPage