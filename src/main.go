package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = "8083"

func main() {
	log.Println("Starting server ...")
	http.HandleFunc("/", visitHandler)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Println("ERROR server:", err)
	}
}

func visitHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "visited %s:%s\n", req.Host, PORT)
}
