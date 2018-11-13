package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	fmt.Printf("hello\n")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context){
		c.JSON(200, gin.H{
			"response": "welcometozombocom",
		})
	})
	r.Run()
}