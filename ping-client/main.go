package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	var server string
	flag.StringVar(&server, "server", "192.168.1.100", "Server running the dump1090-fa program")

	var port int
	flag.IntVar(&port, "port", 8080, "Port to communication to dump1090-fa on")

	var useHttps bool
	flag.BoolVar(&useHttps, "https", false, "Flag to control if HTTPS is used or not")
	var https string = "http"
	if useHttps {
		https = "https"
	}

	flag.Parse()

	var builder strings.Builder
	fmt.Fprintf(&builder, "%s://%s:%d", https, server, port)
	var svr string = builder.String()

	chIsDone := make(chan bool, 1)
	chStopRequested := make(chan bool, 1)

	go loopForever(svr, chIsDone, chStopRequested)

	fmt.Println("Thread is running!")

	// read seems to wait until enter/return is pressed; but once it
	// has been pressed, send a message on the channel that it is time
	// to shut down
	fmt.Print("Press 'enter' to stop...\n\n")
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	chStopRequested <- true

	fmt.Println("Waiting for thread to finish...")
	<-chIsDone
	fmt.Println("All done!")
}

func loopForever(server string, chWhenDone chan<- bool, chStopRequested <-chan bool) {

	fmt.Printf("Starting to ping %s for aircraft data...\n", server)

	chDoneSleeping := make(chan bool, 1)

	isStopRequested := false

	for !isStopRequested {

		sleepAndSignal(chDoneSleeping)

		select {
		// need to check stop requested first; otherwise, it'll check if it's done sleeping
		// first, and then call pong again before checking if a stop has been requested
		case <-chStopRequested:
			fmt.Println("Stop requested!")
			isStopRequested = true

		case <-chDoneSleeping:
			var json string = getAircraft(server)
			fmt.Printf("%v\n", json)
		}
	}
	chWhenDone <- true
}

func sleepAndSignal(onDone chan<- bool) {
	time.Sleep(time.Second)
	onDone <- true
}

func getAircraft(server string) string {

	// https://github.com/adsbxchange/dump1090-fa/blob/master/README-json.md
	resp, err := http.Get(server + "/data/aircraft.json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	return sb
}
