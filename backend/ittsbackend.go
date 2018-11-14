package main

import (
	"github.com/gin-gonic/gin"
	"flag"
)

func main() {
    // create a flag to take in the port number as a CLI flag
	servicePort := flag.String("port", "42069", "target port to run the service")
	flag.Parse()

	router := gin.Default()
	router.GET("/test", test)
	router.GET("/test/:test", testParam)

	// run our router with the specified or default port number
	router.Run(":"+*servicePort)
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
