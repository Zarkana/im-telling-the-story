/*
This is heavily based on "https://github.com/coreos/go-oidc/tree/v2/example/idtoken"
It's worth mentioning naming this "oauth.go" is somewhat of a misnomer, as we don't really do anything with oauth
Also: we may need to change this over to not be "implicit flow" depending on how much we care about security
*/
package auth

import (
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

// bad for obvious reasons
// We use the nonce as something we can put inside the return JWT and this allows us to make sure that nothing terrible happened.
const appNonce = "a super secret nonce"

func init() { // todo: i'm pretty sure all of this isn't a great way to do this. I believe we should be doing all this every time anyone tries to login, not just once when we start the server.
	// read our secrets
	secrets := ReadJSON()

	clientID = secrets.Web.ClientID
	clientSecret := secrets.Web.ClientSecret

	// Disclosure: I'm not sure how you use context or if this is how it is
	ctx = context.Background()
	// (it isn't)

	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}

	// config that oauth2
	config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		// This is something we need to change when we move to a real platform
		RedirectURL: "http://127.0.0.1:5555/auth/google/callback",
		Scopes:      []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// we need to fix this, but not really sure how lmao
	// I read more about this. we can use the state to sort of verify nothing bad happened, but we could also use it to redirect people to the right place.
	state = "foobar" // Don't do this in production.

}

// Routes is all the routes we need for google OIDC auth
func Routes(route *gin.RouterGroup) {
	google := route.Group("/google")
	{
		// this one is pretty easy just redirecting to google authentication with our state variable
		google.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusFound, config.AuthCodeURL(state, oidc.Nonce(appNonce)))
		})
		google.GET("/callback", googleCallback)
	}
}

// this is how we handle the callback from google
// This function is at the risk of becoming mega spaghetti
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

	if idToken.Nonce != appNonce {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid Nonce"))
	}

	// formatting our response to print out. We're going to want to do something with this lol
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *IDTokenClaims // ID Token payload is just JSON.
	}{oauth2Token, new(IDTokenClaims)}

	// this is unmarshalling the claims from the idToken (stuff like name, email etc)
	// it's just json.RawMessage so just a string :v
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// This sets our cookie so we can make some reasonable claim that we know who they are. We might want to use our own user ID for this not google's sub.
	SetCookie(c, GetSignedToken(string(resp.IDTokenClaims.Sub)))
	c.JSON(200, resp)
}
