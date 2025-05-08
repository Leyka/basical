package event

import (
	"strings"
	"time"

	"github.com/leyka/basical/internal/config"
)

type EventService struct {
	config *config.Config
}

func NewEventService(config *config.Config) *EventService {
	return &EventService{
		config: config,
	}
}

func FilterEvents(events []Event, keywords []string) []Event {
	var filteredEvents []Event
	for _, e := range events {
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(e.Title), strings.ToLower(keyword)) {
				filteredEvents = append(filteredEvents, e)
				break
			}
		}
	}

	return filteredEvents
}

func (s *EventService) GetTomorrowEvents(keywords []string) ([]Event, error) {
	events, err := fetchEvents(s.config)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, nil
	}

	filteredEvents := FilterEvents(events, keywords)
	if len(filteredEvents) == 0 {
		return nil, nil
	}

	var tomorrowEvents []Event
	tomorrow := time.Now().AddDate(0, 0, 1)
	year, month, day := tomorrow.Date()

	for _, e := range filteredEvents {
		eYear, eMonth, eDay := e.Date.Date()
		if year == eYear && month == eMonth && day == eDay {
			tomorrowEvents = append(tomorrowEvents, e)
		}
	}

	if len(tomorrowEvents) == 0 {
		return nil, nil
	}

	return tomorrowEvents, nil
}
