package kolgram

type TelegramInterface interface {
	New(token string) Telegram

	get(method string) map[string]interface{}
	post(method string) map[string]interface{}
}
