package main

import (
	"github.com/gocolly/colly"
	"github.com/kr/pretty"
)

type details struct {
	header string
	desc   string
	href   string
}

func main() {
	c := colly.NewCollector()

	allitems := []details{}

	// Find and visit all links
	c.OnHTML("div.three div.row div.column", func(e *colly.HTMLElement) {
		items := details{
			header: e.ChildText("h2.ui"),
			desc:   e.ChildText("p"),
			href:   e.ChildAttr("a", "href"),
		}
		allitems = append(allitems, items)

	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("http://go-colly.org/")
	pretty.Print(allitems)
}
