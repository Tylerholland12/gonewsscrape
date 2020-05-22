package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

const site = "https://twitter.com/i/events/1249735459527254017"

type tweetnews struct {
	Name     string
	Username string
	Message  string
}

func main() {
	col := colly.NewCollector()

	messages := []tweetnews{}

	col.OnHTML(".tweetnews", func(e *colly.HTMLElement) {
		messages = append(messages, tweetnews{
			Name:     e.ChildText(".account-group .fullname"),
			Username: e.ChildText(".account-group .username"),
			Message:  e.ChildText(".tweetnews-text"),
		})
	})

	err := col.Visit(site)
	if err != nil {
		panic(err)
	}

	col.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("number of tweets:", len(messages))
}
