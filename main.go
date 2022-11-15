package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ishant-tata/NetFlix_Project_MongoDB/router"
)

func main() {
	fmt.Println("Welcome to NetFlix Project (GoLang+ MongoDB)")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8010", r))
	fmt.Println("Listening on port 8010...")
}
