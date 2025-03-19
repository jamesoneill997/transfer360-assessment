package publish

import (
	"log"
)

func ListenForMessages(publisher *MockPublisher) {
	for msg := range publisher.Messages {
		log.Printf("Published message for Vehicle: %v", msg.Vrm)
	}
}
