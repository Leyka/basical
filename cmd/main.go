package main

import (
	"strings"

	"github.com/leyka/basical/internal/config"
	"github.com/leyka/basical/internal/event"
	"github.com/leyka/basical/internal/log"
	"github.com/leyka/basical/internal/notifier"
	"github.com/leyka/basical/internal/utils"
)

var logger = log.GetLogger()

func getTomorrowCollectsEvents(config *config.Config) ([]event.Event, error) {
	eventService := event.NewEventService(config)

	keywords := []string{"collecte"}
	events, err := eventService.GetTomorrowEvents(keywords)
	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return nil, nil
	}

	return events, nil
}

func main() {
	defer log.LogFile.Close()

	config, err := config.LoadConfig("config.yml")
	if err != nil {
		logger.Fatal("Error loading site config", "error", err)
	}

	tomorrowEvents, err := getTomorrowCollectsEvents(config)
	if err != nil {
		logger.Fatal("Error fetching tomorrow's events", "error", err)
	}

	if len(tomorrowEvents) == 0 {
		logger.Info("No events found for tomorrow")
		return
	}

	telegram := notifier.NewTelegram(config.Telegram.BotToken)
	for _, e := range tomorrowEvents {
		formattedDate := strings.ToLower(utils.GetFormattedDateInFrench(e.Date))
		msg := "*" + e.Title + "*\n" + "Demain le " + formattedDate + "\n" + e.Link

		for _, chatID := range config.Telegram.ChatIDs {
			err := telegram.SendMessage(chatID, msg)
			if err != nil {
				logger.Error("Error sending message to Telegram", "error", err, "chat_id", chatID)
			} else {
				logger.Info("Message sent to Telegram", "chat_id", chatID)
			}
		}
	}
}
