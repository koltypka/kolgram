package main

import (
	"encoding/json"
	"fmt"

	"github.com/koltypka/kolgram/request"
)

//"kolgram";
//"SQL";
//"yamlReader";
//"logic"

func main() {
	Request := request.New("https://api.hh.ru") //https://api.gismeteo.net/v2/weather/current/?latitude=54.35&longitude=52.52

	Request.AddParam("employer_id", "1975264")
	Request.AddParam("page", "2")

	aaa, _ := Request.Get("vacancies/")

	var hm interface{}

	json.Unmarshal(aaa, &hm)

	myMap := hm.(map[string]interface{})

	items := myMap["items"].([]interface{})

	//fmt.Println(items)
	for _, value := range items {
		fmt.Println(value.(map[string]interface{})["id"])
	}

	// fmt.Print(bb)

	//parametrs := yamlReader.GetYamlFile('parameters.yaml')

	//connection := SQL.GetConnection()

	//EvaGoer := kolgram.GetConnection(parametrs['token'])

}
