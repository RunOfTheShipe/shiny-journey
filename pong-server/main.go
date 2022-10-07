package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pong", getPong)

	router.Run("localhost:6543")
}

func getPong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Pong")
}
