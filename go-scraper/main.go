package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

type grammar struct {
	Id         string `json:"id"`
	Eng        string `json:"eng"`
	Jp         string `json:"jp"`
	Definition string `json:"def"`
	Link       string `json:"link"`
}

func generateId(id string) string {
	newId := strings.TrimSpace(id)
	newId = strings.ReplaceAll(newId, " ", "-")
	newId = strings.ToLower(newId)
	return newId
}

func main() {
	visitedLinks := make(map[string]string)
	var linksToVisit []string
	var grammarList []interface{}

	c := colly.NewCollector(colly.AllowedDomains("jlptsensei.com"))
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})
	c.OnHTML(".jl-row", func(e *colly.HTMLElement) {
		id := generateId(e.ChildText(".jl-td-gr"))
		newGrammar := grammar{
			Id:         id,
			Eng:        e.ChildText(".jl-td-gr"),
			Jp:         e.ChildText(".jl-td-gj"),
			Definition: e.ChildText(".jl-td-gm"),
			Link:       e.ChildAttr("td > .jl-link", "href"),
		}
		var grammarTupel []interface{}
		grammarTupel = append(grammarTupel, newGrammar.Id)
		grammarTupel = append(grammarTupel, newGrammar)
		grammarList = append(grammarList, grammarTupel)
	})
	c.OnHTML(".page-numbers", func(e *colly.HTMLElement) {
		if e.Attr("href") != "" {
			if _, ok := visitedLinks[e.Attr("href")]; !ok {
				linksToVisit = append(linksToVisit, e.Attr("href"))
				fmt.Println("adding to list: ", e.Attr("href"))
			}
		}
	})
	c.Visit("https://jlptsensei.com/jlpt-n5-grammar-list/page/1/")
	visitedLinks["https://jlptsensei.com/jlpt-n5-grammar-list/page/1/"] = "start"

	for len(linksToVisit) > 0 {
		link := linksToVisit[0]
		linksToVisit = append(linksToVisit[:0], linksToVisit[1:]...)
		visitedLinks[link] = "visited"
		c.Visit(link)
	}

	// time to output the grammar!
	file, _ := json.MarshalIndent(grammarList, "", "\t")
	fmt.Println(string(file))

	_ = ioutil.WriteFile("test.json", file, 0644)
}
