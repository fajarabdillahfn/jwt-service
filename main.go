package main

import (
	"jwt-service/database"
	"jwt-service/routers"
	"log"
	"os"
)

func main() {
	database.StartDB()

	port := os.Getenv("PORT")

	r := routers.StartApp()
	log.Println("starting app...")
	r.Run(":" + port)
}
