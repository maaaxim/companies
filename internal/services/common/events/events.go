package events

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Event interface {
	GetName() string
	Marshal() ([]byte, error)
}

//go:generate mockgen -source=$GOFILE -destination=mocks/$GOFILE
type EventsPublisher interface {
	GoPublishEvent(event Event)
}

type blankEvent struct {
	Name     string `json:"name"`
	Uuid     string `json:"uuid"`
	EventDto any    `json:"properties"`
}

func newBlankEvent(name string, event any) blankEvent {
	return blankEvent{
		Name:     name,
		Uuid:     uuid.New().String(),
		EventDto: event,
	}
}

func (e blankEvent) GetName() string {
	return e.Name
}

func (e blankEvent) Marshal() ([]byte, error) {
	return json.Marshal(e) //nolint:wrapcheck
}
