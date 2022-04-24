package main

import (
	"fmt"

	"github.com/Samaraferreira/go-api-rest/database"
	"github.com/Samaraferreira/go-api-rest/routes"
)

func main() {
	database.Connect()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
