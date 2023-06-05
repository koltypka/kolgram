package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/koltypka/kolgram/request"
)

//"kolgram";
//"SQL";
//"yamlReader";
//"logic"

func main() {
	Request := request.New("api.hh.ru") //https://api.gismeteo.net/v2/weather/current/?latitude=54.35&longitude=52.52
	v := url.Values{}
	/*v.Add("latitude", "54.35")
	v.Add("longitude", "52.52")*/

	aaa, _ := Request.Run("vacancies/", v)

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
