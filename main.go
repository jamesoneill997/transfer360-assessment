package main

import (
	"flag"
	"fmt"
	"log"
	"t360/api/models"
	"t360/api/publish"
	"t360/api/search"
)

func main() {
	vrm := flag.String("vrm", "", "The VRM of the vehicle")
	contraventionDate := flag.String("contravention_date", "", "The contravention date of the vehicle")

	flag.Parse()

	if *vrm == "" || *contraventionDate == "" {
		log.Fatal("Both 'vrm' and 'contravention_date' are required")
	}

	publisher := publish.NewMockPublisher()

	lookup := models.Lookup{
		Vrm:               *vrm,
		ContraventionDate: *contraventionDate,
	}

	search.Search([]models.Lookup{lookup}, publisher)

	fmt.Println("Search completed successfully.")
}
