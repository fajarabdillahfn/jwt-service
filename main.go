package main

import (
	"jwt-service/database"
	"jwt-service/routers"
	"log"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	log.Println("starting app...")
	r.Run(":8080")
}
