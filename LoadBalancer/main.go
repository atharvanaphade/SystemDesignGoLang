package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var port string = os.Getenv("PORT_FIRST")

func main() {
	fmt.Println(port)
	router := gin.Default()
	router.GET("/", hello)

	router.Run("0.0.0.0:" + string(port))
}

func hello(c *gin.Context) {
	fmt.Println("Port: " + port)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
