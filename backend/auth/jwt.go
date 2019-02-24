/*
   file with a bunch of nice little helpers for cookies and stuff
*/
package auth

import (
	"encoding/json"
	"fmt"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
)

// don't do this obv
var key = []byte("What the fuck did you just fucking say about me, you little bitch?")

// getSignedToken allows you to pass a string to be encoded and return a JWT signed with our key
func getSignedToken(json string) string {

	token, err := jose.Sign(json, jose.HS256, key)

	if err == nil {
		fmt.Println("Token ok")
		return token
	}
	return ""
}

// verifyToken Allows you to pass a JWT and it will return whether it could be verified and what the string is
func verifyToken(jwt string) (JWTFormat, bool) {
	payload, _, err := jose.Decode(jwt, key)
	if err != nil {
		return *new(JWTFormat), false
	}
	jwtStruct := new(JWTFormat)
	json.Unmarshal([]byte(payload), &jwtStruct)
	return *jwtStruct, true
}

// SetCookie will when given a string of text and context, will allow you to set a cookie named "JWT" for 127.0.0.1. Change if not that lmao
func SetCookie(c *gin.Context, text JWTFormat) {
	fmt.Println(text)
	val, err := json.Marshal(text)
	if err != nil {
		panic(err)
	}
	c.SetCookie("JWT", getSignedToken(string(val)), 3600, "/", "127.0.0.1", false, true)
}

// GetCookie allows you to return the value of the JWT token
// and whether the value could be verified
// returns "", false if no cookie
func GetCookie(c *gin.Context) (JWTFormat, bool) {
	cookie, err := c.Cookie("JWT")
	if err != nil {
		fmt.Printf(err.Error())
		return *new(JWTFormat), false
	}
	return verifyToken(cookie)
}

// GetUserID takes a context and returns the Userid of the current user
func GetUserID(c *gin.Context) int64 {
	jwt, verified := GetCookie(c)
	if !verified {
		return 0
	}
	return jwt.UserID
}
