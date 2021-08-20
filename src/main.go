package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT = "8083"

func main() {
	log.Println("Starting server ...")
	http.HandleFunc("/", visitHandler)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Println("ERROR server:", err)
	}
}

// visitHandler returns the number of times a request has been made to the server
func visitHandler(res http.ResponseWriter, req *http.Request) {
	WriteLog(req.Host)
	count := VisitCount(req.Host)
	fmt.Fprintf(res,
		"You are talking to instance %s.\nThis is request %[2]d to this instance and request %[2]d to the cluster.\n",
		req.Host, count)
}

// WriteLogs writes a line to instance log file when visited.
func WriteLog(instance string) {

	ilog, err := os.OpenFile(instance, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	defer ilog.Close()
	if err != nil {
		log.Println("ERROR opening instance log file", err)
	}

	log.SetOutput(ilog)
	log.Println("instance visited")

}

//  VisitCount returns the number of logs (visits) written to instance log file.
func VisitCount(instance string) int {
	var count int

	ilog, err := os.Open(instance)
	defer ilog.Close()
	if err != nil {
		log.Println("ERROR opening instance log file", err)
	}

	scanner := bufio.NewScanner(ilog)
	for scanner.Scan() {
		count++
	}

	return count
}
