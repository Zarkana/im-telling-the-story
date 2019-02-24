package main

import (
	"flag"
	"fmt"

	Auth "./auth"
	DB "./db"
	"github.com/gin-gonic/gin"
)

func main() {
	// create a flag to take in the port number as a CLI flag
	servicePort := flag.String("port", "5555", "target port to run the service")
	flag.Parse()

	router := gin.Default()
	router.GET("/test", test)
	router.GET("/test/:test", testParam)
	// we make a route to pass down into our auth
	auth := router.Group("/auth")
	Auth.Routes(auth)
	authorized := router.Group("/secret")
	authorized.Use(Auth.LoginMiddleware())
	authorized.GET("/loggedin", loggedIn)
	// testing our database functions
	DB.Test()
	fmt.Println(Auth.ReadJSON())
	// run our router with the specified or default port number
	router.Run(":" + *servicePort)
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

func loggedIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"username": DB.GetScreenName(Auth.GetUserID(c)),
	})
}
