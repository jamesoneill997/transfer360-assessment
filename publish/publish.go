package publish

import (
	"context"
	"log"
	"t360/api/models"
)

// MockPublisher simulates Pub/Sub
type MockPublisher struct {
	Messages chan models.LookupResult
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{Messages: make(chan models.LookupResult, 100)}
}

func (m *MockPublisher) Publish(ctx context.Context, msg models.LookupResult) error {
	log.Printf("Mock publish: %v", msg)
	m.Messages <- msg
	return nil
}
