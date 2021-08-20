package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT = "8083"
var DBDIR = "/db/"

func main() {
	log.Println("Starting server ...")
	http.HandleFunc("/", visitHandler)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Println("ERROR server:", err)
	}
}

// visitHandler returns the number of times a request has been made to the server
func visitHandler(res http.ResponseWriter, req *http.Request) {
	// set instance and cluster names
	instance, err := os.Hostname()
	if err != nil {
		log.Println("ERROR getting instance Hostname", err)
	}
	cluster := req.Host

	WriteLogs(DBDIR+instance, DBDIR+cluster)

	icount, ccount := VisitCounts(DBDIR+instance, DBDIR+cluster)

	fmt.Fprintf(res,
		"You are talking to instance %s:%s.\nThis is request %d to this instance and request %d to the cluster.\n",
		instance, PORT, icount, ccount)
}

// WriteLogs writes a line to instance log file when visited.
func WriteLogs(instance, cluster string) {
	// instance log
	ilog, err := os.OpenFile(instance, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	defer ilog.Close()
	if err != nil {
		log.Println("ERROR opening instance log file", err)
	}
	log.SetOutput(ilog)
	log.Println("instance visited")

	// cluster log
	clog, err := os.OpenFile(cluster, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	defer clog.Close()
	if err != nil {
		log.Println("ERROR opening cluster log file", err)
	}
	log.SetOutput(clog)
	log.Println("cluster visited")

}

//  VisitCount returns the number of logs (visits) written to instance log file.
func VisitCounts(instance, cluster string) (int, int) {
	var icount, ccount int

	// instance log counts
	ilog, err := os.Open(instance)
	defer ilog.Close()
	if err != nil {
		log.Println("ERROR opening instance log file", err)
	}

	scanner := bufio.NewScanner(ilog)
	for scanner.Scan() {
		icount++
	}

	// cluster log counts
	clog, err := os.Open(cluster)
	defer clog.Close()
	if err != nil {
		log.Println("ERROR opening cluster log file", err)
	}

	scanner = bufio.NewScanner(clog)
	for scanner.Scan() {
		ccount++
	}

	return icount, ccount
}
