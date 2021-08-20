package main

import (
	"bufio"
	"log"
	"os"
)

// WriteLogs writes a line to instance and cluster log files when visited.
func WriteLogs(instance, cluster string) {

	// instance logging
	ilog, err := os.OpenFile(instance, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	defer ilog.Close()
	check("ERROR opening instance log file", err)

	log.SetOutput(ilog)
	log.Println("instance visited")

	// cluster logging
	clog, err := os.OpenFile(cluster, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	defer clog.Close()
	check("ERROR opening cluster log file", err)

	log.SetOutput(clog)
	log.Println("cluster visited")

}

// VisitCounts returns the number of visits to instance and cluster
// Visits are calculated from the number of lines written to log files.
func VisitCounts(instance, cluster string) (int, int) {
	var icount, ccount int

	// instance log counts
	ilog, err := os.Open(instance)
	defer ilog.Close()
	check("ERROR opening instance log file", err)

	scanner := bufio.NewScanner(ilog)
	for scanner.Scan() {
		icount++
	}

	// cluster log counts
	clog, err := os.Open(cluster)
	defer clog.Close()
	check("ERROR opening cluster log file", err)

	scanner = bufio.NewScanner(clog)
	for scanner.Scan() {
		ccount++
	}

	return icount, ccount
}

// check handles the error checking only for readibility purposes
func check(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
