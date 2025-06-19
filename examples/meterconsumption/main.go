package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"octogo"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, falling back to system environment variables")
	}

	apiKey := os.Getenv("OCTOPUS_API_KEY")
	if apiKey == "" {
		log.Fatal("OCTOPUS_API_KEY environment variable is required")
	}

	mpan := os.Getenv("ELECTRIC_MPAN")
	if mpan == "" {
		log.Fatal("ELECTRIC_MPAN environment variable is required")
	}

	serialNumber := os.Getenv("ELECTRIC_SERIAL_NUMBER")
	if serialNumber == "" {
		log.Fatal("ELECTRIC_SERIAL_NUMBER environment variable is required")
	}

	client := octogo.NewClient(apiKey)
	ctx := context.Background()

	fmt.Printf("Fetching consumption details")

	order := "-period"
	page_size := 10
	consumption, resp, err := client.Meter.GetConsumption(ctx, mpan, serialNumber, &octogo.ConsumptionOptions{
		OrderBy:  &order,
		PageSize: &page_size,
	})
	if err != nil {
		log.Fatalf("Error fetching consumption details: %v", err)
	}

	fmt.Printf("HTTP Status: %s\n", resp.Status)
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Consumption Details:\n")
		fmt.Printf("  Total Readings: %d\n", consumption.Count)
		if len(consumption.Results) > 0 {
			for i := 0; i < len(consumption.Results); i++ {
				reading := consumption.Results[i]
				fmt.Printf("    %d. Consumption: %.3f kWh, Period: %s to %s\n",
					i+1, reading.Consumption,
					reading.IntervalStart.Format("2006-01-02 15:04:05"),
					reading.IntervalEnd.Format("2006-01-02 15:04:05"))
			}
		}
	} else {
		fmt.Printf("API returned status code: %d\n", resp.StatusCode)
	}
}
