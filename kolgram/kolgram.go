package kolgram

import (
	"fmt"

	request "github.com/koltypka/kolRequest/kolRequest"
)

type Telegram struct {
	Request request.Request
}

func New(token string) Telegram {
	return Telegram{request.New("https://api.telegram.org/bot" + token)}
}

func (Telegram *Telegram) GetUpdates() {
	result, _ := Telegram.Request.Post("/getUpdates")

	fmt.Println(result.ToJson())
}
