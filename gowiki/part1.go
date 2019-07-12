package main

import (
	"fmt"
	"io/ioutil"
)

// describes how page data will be stored in memory
type Page struct {
	Title string
	Body []byte
}

// for persistent storage
// save the page's Body to a text file.
// For simplicity, we will use the Title as the file name
// succeed return nil, 0600 - current user only
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// constructs the file name from the title param, 
// reads the file's content a new variable body
// returns a pointer to a Page literal constructed with the proper title and  body values
/* notes: 
 *  functions can return multiple values. io.ReadFile return []byte and error
 *  in loadPage, error isn't being handled yet, 
 *  " blank identifier" (_) is used to throw away the error return value
 * 
 * func loadpage(title string) *Page {
 * 	filename := title + ".txt"
 * 	body, _ := ioutil.ReadFile(filename)
 * 	return &Page{Title: title, Body: body}
 * }
 * 
 */
// if ReadFile encounter error, the file not exist
func loadPage(title string) (*Page, error) {
	filename :=title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// testing
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a simple Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}