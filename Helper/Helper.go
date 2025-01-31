package helper

import (
	"encoding/csv"
	"fmt"
	"github.com/joho/godotenv"
	typ "github.com/joy-adhikaryy/Go-Scraper/Types"
	"log"
	"os"
	"path/filepath"
)

func GetEnv(key string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		log.Printf("Warning: Environment variable '%s' not found", key)
	}
	return value
}

//func DumpToCSV(stock []typ.Stocks{}) {
//	file, err := os.Create(filepath.Join("./Data", "stocks.csv"))
//	if err != nil {
//		log.Fatalln("Failed to create output CSV file", err)
//	}
//	defer file.Close()
//
//	writer := csv.NewWriter(file)
//	headers := []string{
//		"Company",
//		"Price",
//		"Change",
//		"Market Status",
//	}
//	writer.Write(headers)
//
//	for _, stock := range stocks {
//		record := []string{
//			stock.Company,
//			stock.Price,
//			stock.Change,
//			stock.MarketStatus,
//		}
//		writer.Write(record)
//	}
//	defer writer.Flush()
//}

func DumpToCSV(stocks []typ.Stocks) error {
	// Ensure data directory exists
	err := os.MkdirAll("./Data", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	// Create file with proper permissions
	file, err := os.Create(filepath.Join("./Data", "stocks.csv"))
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"Company",
		"Price",
		"Change",
		"Market Status",
	}

	// Write headers
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	// Write records
	for _, stock := range stocks {
		record := []string{
			stock.Company,
			stock.Price,
			stock.Change,
			stock.MarketStatus,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write stock record: %w", err)
		}
	}

	return nil
}
