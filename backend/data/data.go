package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/kr/pretty"
)

type Details struct {
	Header string `json:"header"`
	Desc   string `json:"desc"`
	Href   string `json:"href"`
}

func FetchData(url string, result chan<- []Details) {

	allitems := []Details{}
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("body", func(e *colly.HTMLElement) {

		items := Details{
			Header: e.ChildText("h2"),
			Desc:   e.ChildText("p"),
			Href:   e.DOM.Find("a[href]").AttrOr("href", ""),
		}
		allitems = append(allitems, items)
	})

	c.Visit(url)
	pretty.Println(allitems)

	// Marshal the struct array into JSON bytes
	jsonData, err := json.Marshal(allitems)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	result <- []Details(allitems)

	// Write JSON data to a file
	err = os.WriteFile("products.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to products.json successfully")

	return
}
