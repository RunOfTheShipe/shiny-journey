package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pong", getPong)
	router.POST("/aircraft", processAircraft)

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

	log.Printf("Process request: timestamp=%f messages=%d aircraft=%d\n", request.Now, request.Messages, len(request.Aircraft))

	c.IndentedJSON(http.StatusOK, "Processed")
}
