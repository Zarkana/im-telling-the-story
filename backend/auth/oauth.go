/*
This is heavily based on "https://github.com/coreos/go-oidc/tree/v2/example/idtoken"
*/
package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	oidc "github.com/coreos/go-oidc"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	ctx      context.Context
	config   oauth2.Config
	state    string
	clientID string
)

func init() {

	// read our secrets
	secrets := ReadJSON()

	clientID = secrets.Web.ClientID
	clientSecret := secrets.Web.ClientSecret

	// Disclosure: I'm not sure how you use context or if this is how it is
	ctx = context.Background()

	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}

	// config that oauth2
	config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:5555/auth/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// we need to fix this, but not really sure how lmao
	state = "foobar" // Don't do this in production.

}

// Routes is all the routes we need for google OIDC auth
func Routes(route *gin.RouterGroup) {
	google := route.Group("/google")
	{
		// this one is pretty easy just redirecting to google authentication with our state variable
		google.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusFound, config.AuthCodeURL(state))
		})
		google.GET("/callback", googleCallback)
	}
}

// this is how we handle the callback from google
func googleCallback(c *gin.Context) {

	// our OpenID Connect handler
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	// The state should be the same as when we sent it.
	// I'm not terribly sure how this makes it more secure but trust it do
	if c.Request.URL.Query().Get("state") != state {
		c.AbortWithError(400, errors.New("state did not match"))
		return
	}
	// we get the oauth token
	oauth2Token, err := config.Exchange(ctx, c.Request.URL.Query().Get("code"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to exchange token: "+err.Error()))
		return
	}
	// the IDToken is the thing we're looking for, this identifies the user to us
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, errors.New("No id_token field in oauth2 token"))
		return
	}
	// verify the token
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to verify ID Token: "+err.Error()))
		return
	}

	// formatting our response to print out. We're going to want to do something with this lol
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	// this is unmarshalling the claims from the idToken (stuff like name, email etc)
	// it's just json.RawMessage so just a string :v
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// formatting
	// data, err := json.MarshalIndent(resp, "", "    ")
	// if err != nil {
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }
	c.JSON(200, resp)
}
