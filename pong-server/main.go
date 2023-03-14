package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "embed"

	"github.com/gin-gonic/gin"
)

var messageCount int = 0
var aircraftNum int = 0

//go:embed index.html
var indexHtml []byte

func main() {
	router := gin.Default()
	router.GET("/pong", getPong)
	router.POST("/aircraft", processAircraft)
	router.GET("/stats", getStats)
	router.GET("/", getIndex)

	router.Run("localhost:6543")
}

func getPong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Pong")
}

func processAircraft(c *gin.Context) {

	var request AircraftRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("Bad JSON!")
		return
	}

	messageCount = request.Messages
	aircraftNum = len(request.Aircraft)

	log.Printf("Process request: timestamp=%f messages=%d aircraft=%d\n", request.Now, messageCount, aircraftNum)

	c.IndentedJSON(http.StatusOK, "Processed")
}

func getIndex(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(indexHtml)
}

// streaming example is from: https://github.com/gurleensethi/go-server-sent-event-example/tree/main
func getStats(c *gin.Context) {
	request := c.Request
	writer := c.Writer

	writer.Header().Set("Content-Type", "text/event-stream")

	// as long as the request is held open, watch the variables and post
	// an update every second
	dataCh := make(chan MetaData)
	go watchStats(request.Context(), dataCh)

	for data := range dataCh {
		serverEvent, err := formatServerEvent("metadata", data)
		if err != nil {
			fmt.Println(err)
			break
		}

		_, err = fmt.Fprintf(writer, serverEvent)
		if err != nil {
			fmt.Println(err)
			break
		}

		writer.Flush()
	}

	log.Printf("getStats() Done!")

}

func watchStats(ctx context.Context, dataCh chan<- MetaData) {

	ticker := time.NewTicker(time.Second)

	var keepGoing bool = true

	for keepGoing {
		select {
		case <-ctx.Done():
			keepGoing = false
		case <-ticker.C:
			dataCh <- MetaData{messageCount, aircraftNum}
		}
	}

	ticker.Stop()
	close(dataCh)
	log.Printf("watchStats() Done!")
}

func formatServerEvent(eventName string, data MetaData) (string, error) {
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("event: %s\n", eventName))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", buff.String()))

	return sb.String(), nil
}
