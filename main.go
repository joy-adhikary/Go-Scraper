package main

import (
	"fmt"
	"github.com/gocolly/colly"
	hlp "github.com/joy-adhikaryy/Go-Scraper/Helper"
	"strings"
	"time"
)

func main() {

	url := hlp.GetEnv("URLCA")

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.AllowURLRevisit = true

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
		fmt.Println("Currently Visiting", r.URL)
	})

	c.OnHTML("div.content-wrapper", func(e *colly.HTMLElement) {
		e.ForEach("div.container table", func(_ int, table *colly.HTMLElement) {
			// getting Header
			header := table.ChildText("thead tr:first-child th")

			if strings.Contains(header, "Interbank USD/BDT") {
				// Table A (Interbank rates)
				cleanedHeader := strings.ReplaceAll(header, "\n", "")
				cleanedHeader = strings.ReplaceAll(cleanedHeader, "\t", "")
				fmt.Println(cleanedHeader)

				table.ForEach("tbody tr", func(_ int, row *colly.HTMLElement) {
					currency := row.ChildText("td:nth-child(1)")
					dayLow := row.ChildText("td:nth-child(2)")
					dayHigh := row.ChildText("td:nth-child(3)")
					currentWar := row.ChildText("td:nth-child(4)")

					fmt.Printf("%s: DayLow=%s, DayHigh=%s, CurrentWar=%s\n",
						currency, dayLow, dayHigh, currentWar)
				})

			} else if strings.Contains(header, "Cross rates") {
				// Table B (Cross rates)
				date := table.ChildText("thead tr:first-child th")
				fmt.Printf("\nCross Rates (%s):\n", strings.TrimSpace(date))

				table.ForEach("tbody tr", func(_ int, row *colly.HTMLElement) {
					currency := row.ChildText("td:nth-child(1)")
					buyingRate := row.ChildText("td:nth-child(2)")
					sellingRate := row.ChildText("td:nth-child(3)")

					fmt.Printf("%s: Buy=%s, Sell=%s\n",
						currency, buyingRate, sellingRate)
				})
			}
		})
	})

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	counter := 0

	for range ticker.C {
		fmt.Println("hit", counter)
		c.Visit(url)
		c.Wait()
		counter++
	}
}
