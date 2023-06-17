package main

import (
	"fmt"

	"github.com/quart3z/check-list/api/routes"
	"github.com/quart3z/check-list/internal/database"
)

func main() {

	db := database.ConnectEstablish()

	routes.Routes(db)

	fmt.Println("Server started!")

}
