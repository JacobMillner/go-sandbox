package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("jlptsensei.com"))
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})
	c.OnHTML(".jl-row", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML(".page-numbers", func(e *colly.HTMLElement) {
		link := e.ChildAttr("href", "href")
		link2 := e.ChildAttr("a[href]", "a[href]")
		fmt.Println(e)
		fmt.Println(e.Text)
		fmt.Println(link)
		fmt.Println(link2)
	})
	c.Visit("https://jlptsensei.com/jlpt-n5-grammar-list/")
}
