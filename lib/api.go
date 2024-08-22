package lib

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CallAPI sends a GET request to the specified REST API endpoint and returns the JSON response as a map.
func CallAPI(rest_url string) map[string]interface{} {
	// Declare a variable to hold the JSON response as a map
	var dataMap map[string]interface{}
	// Construct the full URL for the API request using the provided endpoint
	url := "https://" + Rapidapi_host + "/" + rest_url

	// Create a new GET request with the constructed URL
	req, _ := http.NewRequest("GET", url, nil)

	// Add necessary headers to the request
	req.Header.Add("x-rapidapi-key", ApiKey)
	req.Header.Add("x-rapidapi-host", Rapidapi_host)

	// Send the GET request and receive the response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Konverta does not get any response:", err)
	}
	defer res.Body.Close() // Ensure the response body is closed after reading
	// Read the response body
	body, _ := io.ReadAll(res.Body)
	// Unmarshal the JSON response into the dataMap
	err = json.Unmarshal(body, &dataMap)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	// Return the parsed JSON response as a map
	return dataMap
}
