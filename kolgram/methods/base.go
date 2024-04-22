package methods

import (
	request "github.com/koltypka/kolRequest/kolRequest"
	requestResult "github.com/koltypka/kolRequest/kolRequest/result"
)

type Telegram struct {
	Request request.Request
}

func New(token string) *Telegram {
	return &Telegram{request.New("https://api.telegram.org/bot" + token)}
}
func (Telegram *Telegram) AddParam(key, value string) {
	Telegram.Request.AddParam(key, value)
}

func (Telegram *Telegram) Get(method string) (*requestResult.ResultHandler, error) {
	getResult, err := Telegram.Request.Get(method)

	return getResult, err
}

func (Telegram *Telegram) Post(method string) map[string]interface{} {
	getResult, _ := Telegram.Request.Post(method)
	result, _ := getResult.ToJson()

	return result
}
