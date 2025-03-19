package models

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type CompanySeed struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (cs *CompanySeed) VehicleLookup(vehicle Lookup) VehicleSearch {
	apiBase := "https://sandbox-update.transfer360.dev/test_search/"
	apiUrl := apiBase + cs.Namespace

	vehicleSearch := VehicleSearch{
		Reference: uuid.New(),
		Lookup:    vehicle,
	}

	var searchResult LookupResult
	jsonBody, err := json.Marshal(vehicle)
	if err != nil {
		log.Fatalf("Error during json marshal: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, apiUrl, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &searchResult)

	vehicleSearch.SearchResult = searchResult

	return vehicleSearch
}

type VehicleSearch struct {
	Reference    uuid.UUID    `json:"uuid"`
	Lookup       Lookup       `json:"lookup"`
	SearchResult LookupResult `json:"response"`
}

func (vs *VehicleSearch) TagRequest() {
	vs.Reference = uuid.New()
}

type Lookup struct {
	Vrm               string `json:"vrm"`
	ContraventionDate string `json:"contravention_date"`
}

type LookupResult struct {
	Reference         string              `json:"reference"`
	Vrm               string              `json:"vrm"`
	ContraventionDate string              `json:"contravention_date"`
	IsHirerVehicle    bool                `json:"is_hirer_vehicle"`
	LeaseCompany      LeaseCompanyDetails `json:"lease_company"`
}

type LeaseCompanyDetails struct {
	CompanyName  string `json:"companyname"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	AddressLine3 string `json:"addres_line3"`
	AddressLine4 string `json:"address_line4"`
	Postcode     string `json:"postcode"`
}
