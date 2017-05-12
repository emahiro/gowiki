package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var dataDir = "data/"

// Page Title and Body
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	if os.Mkdir(dataDir, 0777) != nil {
		fmt.Println(dataDir + " is already exist")
	}

	filename := p.Title + ".txt"
	err := ioutil.WriteFile(dataDir+filename, p.Body, 0600)
	if err != err {
		fmt.Println("Page not found")
		return err
	}

	return nil
}

func loadPage(title string) (*Page, error) {
	body, err := ioutil.ReadFile(dataDir + title + ".txt")
	if err != nil {
		return nil, err
	}
	page := &Page{Title: title, Body: body}
	return page, nil
}

func main() {
	p := &Page{Title: "test", Body: []byte("This is test")}
	err := p.save()
	if err != nil {
		fmt.Println("Failed")
	}
	page, err := loadPage("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Loading page data is %s", page.Body)
}
