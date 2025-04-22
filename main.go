package main

import (
	"github.com/leonardoavelar/Go-gin-api/database"
	"github.com/leonardoavelar/Go-gin-api/routes"
)

func main() {

	database.Conectar()
	routes.HandlerRequests()

}
