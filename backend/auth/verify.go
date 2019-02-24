package auth

import (
	"errors"

	db "../db"
	"github.com/gin-gonic/gin"
)

// LoginMiddleware is how we check that the user is properly logged in or not
func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, verified := GetCookie(c)
		if !verified {
			c.AbortWithError(401, errors.New("Cookie failure"))
		}
		if db.UserExists(value.UserID) {
			c.Next()
		} else {
			c.Redirect(302, "/auth")
		}
	}
}
