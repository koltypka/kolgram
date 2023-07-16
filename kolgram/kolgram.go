package kolgram

import (
	"encoding/json"
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
	result, _ := Telegram.Request.Get("/getUpdates") //Get("getUpdates")

	var jsonResult interface{}

	json.Unmarshal(result, &jsonResult)

	myMap := jsonResult.(map[string]interface{})
	fmt.Println(myMap)
}
