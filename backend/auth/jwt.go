/*
   file with a bunch of nice little helpers for cookies and stuff
*/
package auth

import (
	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
)

// don't do this obv
var key = []byte("What the fuck did you just fucking say about me, you little bitch?")

// GetSignedToken allows you to pass a string to be encoded and return a JWT signed with our key
func GetSignedToken(json string) string {

	token, err := jose.Sign(json, jose.HS256, key)

	if err == nil {
		return token
	}
	return ""
}

// VerifyToken Allows you to pass a JWT and it will return whether it could be verified and what the string is
func VerifyToken(json string) (string, bool) {
	payload, _, err := jose.Decode(json, key)
	if err != nil {
		return payload, true
	}
	return "", false
}

// SetCookie will when given a string and context, will allow you to set a cookie named "JWT" for 127.0.0.1. Change if not that lmao
func SetCookie(c *gin.Context, token string) {
	c.SetCookie("JWT", token, 3600, "/", "127.0.0.1", false, true)
}

// GetCookie allows you to return the value of the JWT token
func GetCookie(c *gin.Context) string {
	cookie, err := c.Cookie("JWT")
	if err != nil {
		return cookie
	} else {
		return ""
	}
}
