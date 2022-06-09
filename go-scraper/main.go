package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type grammar struct {
	eng        string
	jp         string
	definition string
	link       string
}

func main() {
	visitedLinks := make(map[string]string)
	var linksToVisit []string
	var grammarList []grammar

	c := colly.NewCollector(colly.AllowedDomains("jlptsensei.com"))
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})
	c.OnHTML(".jl-row", func(e *colly.HTMLElement) {
		newGrammar := grammar{
			eng:        e.ChildText(".jl-td-gr"),
			jp:         e.ChildText(".jl-td-gj"),
			definition: e.ChildText(".jl-td-gm"),
			link:       e.ChildAttr("td > .jl-link", "href"),
		}
		grammarList = append(grammarList, newGrammar)
	})
	c.OnHTML(".page-numbers", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
		if e.Attr("href") != "" {
			linksToVisit = append(linksToVisit, e.Attr("href"))
		}
	})
	c.Visit("https://jlptsensei.com/jlpt-n5-grammar-list/")
	visitedLinks["https://jlptsensei.com/jlpt-n5-grammar-list/"] = "start"
}
