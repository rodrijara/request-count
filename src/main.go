// "RequestCount" server counts the number of requests handled by the server
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT = "8083"
var DBDIR = "/db/"

func main() {
	log.Println("Running server...")

	http.HandleFunc("/", visitHandler)

	err := http.ListenAndServe(":"+PORT, nil)
	check("ERROR server:", err)
}

// visitHandler returns the number of times a request has been made to the server.
// The number of times is counted for both instances and cluster.
func visitHandler(res http.ResponseWriter, req *http.Request) {

	// set instance name
	instance, err := os.Hostname()
	check("ERROR getting instance Hostname:", err)

	WriteLogs(DBDIR+instance, DBDIR+"cluster")

	icount, ccount := VisitCounts(DBDIR+instance, DBDIR+"cluster")

	fmt.Fprintf(res,
		"You are talking to instance %s:%s.\n"+
			"This is request %d to this instance and request %d to the cluster.\n",
		instance, PORT, icount, ccount)
}
