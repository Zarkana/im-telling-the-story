package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", test)
	router.GET("/test/:test", testParam)
	router.Run()
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "welcometozombocom",
	})
}

func testParam(c *gin.Context) {
	test := c.Param("test")

	c.JSON(200, gin.H{
		"echo": test,
	})
}
