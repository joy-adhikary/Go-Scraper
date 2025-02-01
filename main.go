package main

import (
	"fmt"
	"github.com/gocolly/colly"
	hlp "github.com/joy-adhikaryy/Go-Scraper/Helper"
	typ "github.com/joy-adhikaryy/Go-Scraper/Types"
)

func main() {

	url := hlp.GetEnv("URLS")
	stocks := []typ.Stocks{}

	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Currently Visiting", r.URL)
	})

	c.OnHTML("section.yf-k4z9w", func(e *colly.HTMLElement) {
		stock := typ.Stocks{}

		stock.Company = e.ChildText("h1")
		stock.Price = e.ChildText("span.base[data-testid='qsp-price']")
		stock.Change = e.ChildText("span.base[data-testid='qsp-price-change']")
		stock.MarketStatus = e.ChildText("span.yf-vednlp")

		stocks = append(stocks, stock)
	})

	c.Wait()

	for _, company := range typ.CompanyNames {
		c.Visit(url + company + "/")
	}

	fmt.Println(stocks)

	hlp.DumpToCSV(stocks)
}
