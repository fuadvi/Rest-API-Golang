package main

import (
	"assement_02/database"
	"assement_02/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
