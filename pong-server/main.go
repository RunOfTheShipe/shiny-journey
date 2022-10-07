package main

import (
	"fmt"
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
		return
	}

	fmt.Printf("Process request: timestamp=%f messages=%d\n", request.Now, request.Messages)
}
