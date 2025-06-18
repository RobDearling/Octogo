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

	client := octogo.NewClient(apiKey)
	ctx := context.Background()

	fmt.Printf("Fetching electricity meter details for MPAN: %s\n", mpan)

	meter, resp, err := client.Meter.GetElectricityMeter(ctx, mpan)
	if err != nil {
		log.Fatalf("Error fetching meter details: %v", err)
	}

	fmt.Printf("HTTP Status: %s\n", resp.Status)

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Meter Details:\n")
		fmt.Printf("  GSP: %s\n", meter.GSP)
		fmt.Printf("  MPAN: %s\n", meter.MPAN)
		fmt.Printf("  Profile Class: %d\n", meter.ProfileClass)
	} else {
		fmt.Printf("API returned status code: %d\n", resp.StatusCode)
	}
}
