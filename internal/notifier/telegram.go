package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Telegram struct {
	ApiURL string
}

func NewTelegram(botToken string) *Telegram {
	return &Telegram{
		ApiURL: "https://api.telegram.org/bot" + botToken + "/sendMessage",
	}
}

type TelegramMessage struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func (t *Telegram) SendMessage(chatId, message string) error {
	telegramMessage := TelegramMessage{
		ChatId:    chatId,
		Text:      message,
		ParseMode: "Markdown",
	}

	jsonData, err := json.Marshal(telegramMessage)
	if err != nil {
		return err
	}

	resp, err := http.Post(t.ApiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}
