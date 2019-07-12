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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title);
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// The function editHandler loads the page 
// (or, if it doesn't exist, create an empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// hard-coded html
	// fmt.Fprintf(w, "<h1>Editing %s</h1>" +
	// 	"<form action=\"/save/%s\" method=\"POST\">" +
	// 	"<textarea name=\"body\">%s</textarea><br>" + 
	// 	"<input type=\"submit\" value=\"Save\">" +
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

// testing
func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http://localhost:8080/view/testPage