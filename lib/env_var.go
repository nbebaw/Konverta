package lib

import (
	"log"
	"os"
)

// getApiKey retrieves the RapidAPI key from the environment variables.
func getApiKey() string {
	// Retrieve the RAPIDAPI_KEY environment variable
	apiKey := os.Getenv("RAPIDAPI_KEY")
	// Check if the environment variable is not set or empty
	if apiKey == "" {
		// Log a fatal error and terminate the program if RAPIDAPI_KEY is not set
		log.Fatal("RAPIDAPI_KEY environment variable not set. Please refer to the documentation.")
	}
	// Return the retrieved API key
	return apiKey
}
