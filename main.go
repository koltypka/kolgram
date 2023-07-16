package main

import (
	"github.com/koltypka/kolgram/kolgram"
)

//"kolgram";
//"SQL";
//"yamlReader";
//"logic"

func main() {
	kolgramOb = kolgram.New("5323142349:AAGUGpI6oMgezVoediL7NYSmn65FVKnb7Q0")
	kolgramOb.getUpdates()
	//getIdhh()
	// fmt.Print(bb)

	//parametrs := yamlReader.GetYamlFile('parameters.yaml')

	//connection := SQL.GetConnection()

	//EvaGoer := kolgram.GetConnection(parametrs['token'])

}
