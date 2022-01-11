package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on port :3000")
	http.ListenAndServe(":3000", http.FileServer(http.Dir("./public")))
}
