package main

import (
	"fmt"
	"github.com/gocolly/colly"
	hlp "github.com/joy-adhikaryy/Go-Scraper/Helper"
	"math/rand"
	"time"
)

func main() {

	url := hlp.GetEnv("URL")

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Currently Visiting", r.URL)
	})

	c.OnHTML("div#matchCenter", func(e *colly.HTMLElement) {

		title := e.ChildText("h1.cb-nav-hdr")
		fmt.Println("Match information:", title)

		score := e.ChildText("div.cb-min-bat-rw")
		scoreData, err := hlp.ParseScoreString(score)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Team: ", scoreData.Team)
		fmt.Println("Runs: ", scoreData.Runs)
		fmt.Println("Wicket: ", scoreData.Wickets)
		fmt.Println("Overs: ", scoreData.Overs)

		status := e.ChildText("div.cb-text-inprogress")
		fmt.Println("Status:", status)
	})

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("hited")
		randomNumber := rand.Intn(100)
		randomNumberStr := fmt.Sprintf("%d", randomNumber)
		c.Visit(url + "?joy=" + randomNumberStr)
	}
}
