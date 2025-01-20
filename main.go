package main

import (
	"CRUD-GIN/config"
	"CRUD-GIN/db"
	"CRUD-GIN/router"
	"log"
	"os"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	if ok := db.Connect(); !ok {
		log.Fatal("error in connecting to DB")
	}
	defer db.DB.Close()

	router := router.RouterSetUp()
	port := os.Getenv("PORT")
	log.Println("server is listening to port: ", port)
	if err := router.Run(); err != nil {
		log.Fatal("error in spining up the server")
	}
}
