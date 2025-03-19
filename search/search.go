package search

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"t360/api/models"
	"t360/api/publish"

	"github.com/google/uuid"
)

func LoadCompanyData() []models.CompanySeed {
	var companies []models.CompanySeed
	b, err := os.ReadFile("./companies.json")
	if err != nil {
		log.Fatalf("Error Loading Seeded Data: %v\n", err)
	}
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&companies)
	return companies
}

func Search(vehicles []models.Lookup, publisher *publish.MockPublisher) []models.VehicleSearch {
	companyData := LoadCompanyData()

	var wg sync.WaitGroup
	resultChan := make(chan models.VehicleSearch, len(companyData)*len(vehicles))

	for _, vehicle := range vehicles {
		for _, company := range companyData {
			wg.Add(1)
			go func(cs models.CompanySeed, v models.Lookup) {
				defer wg.Done()

				apiBase := "https://sandbox-update.transfer360.dev/test_search/"
				apiUrl := apiBase + cs.Namespace
				log.Printf("Starting search for company: %s with API: %s", cs.Name, apiUrl)

				vehicleSearch := models.VehicleSearch{
					Reference: uuid.New(),
					Lookup:    v,
				}

				var searchResult models.LookupResult
				jsonBody, err := json.Marshal(v)
				if err != nil {
					log.Printf("Error during json marshal: %v\n", err)
					resultChan <- vehicleSearch
					return
				}

				bodyReader := bytes.NewReader(jsonBody)
				req, err := http.NewRequest(http.MethodPost, apiUrl, bodyReader)
				if err != nil {
					log.Printf("Error creating request: %v\n", err)
					resultChan <- vehicleSearch
					return
				}

				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					log.Printf("Request failed for company %s: %v", cs.Name, err)
					resultChan <- vehicleSearch
					return
				}
				defer resp.Body.Close()

				log.Printf("Received response for company %s: Status %s", cs.Name, resp.Status)

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Failed to read response for company %s: %v", cs.Name, err)
					resultChan <- vehicleSearch
					return
				}

				log.Printf("Raw response body for company %s: %s", cs.Name, string(body))

				err = json.Unmarshal(body, &searchResult)
				if err != nil {
					log.Printf("Failed to unmarshal response for company %s: %v", cs.Name, err)
					resultChan <- vehicleSearch
					return
				}

				vehicleSearch.SearchResult = searchResult

				// If the search result is a hire vehicle, publish it
				if searchResult.IsHirerVehicle {
					log.Printf("Publishing result for company %s", cs.Name)
					err := publisher.Publish(context.Background(), searchResult)
					if err != nil {
						log.Printf("Failed to publish result for company %s: %v", cs.Name, err)
					}
				}
				resultChan <- vehicleSearch
			}(company, vehicle)
		}
	}

	wg.Wait()
	close(resultChan)

	var allResults []models.VehicleSearch
	for result := range resultChan {
		allResults = append(allResults, result)
	}

	return allResults
}
