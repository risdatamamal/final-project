package main

import (
	"fmt"

	"github.com/risdatamamal/final-project/config"
	"github.com/risdatamamal/final-project/database"
	"github.com/risdatamamal/final-project/router"
)

func main() {
	r := router.StartApp()
	err := database.StartDB()
	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}
	r.Run(config.SERVER_PORT)
}
