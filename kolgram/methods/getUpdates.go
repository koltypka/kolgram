package methods

import (
	"encoding/json"
)

type GetUpdates struct {
	Ok     bool            `json:"ok"`
	Result []UpdatesResult `json:"result"`
}

type UpdatesResult struct {
	UpdateId int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int64  `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type From struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}

func (Telegram *Telegram) GetUpdates() GetUpdates {
	tmpResult, err := Telegram.Get("/getUpdates")
	result := GetUpdates{}

	if err != nil {
		return result
	}

	json.Unmarshal(tmpResult.Body, &result)

	return result
}
