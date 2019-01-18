package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// GoogleJSON is our representation of the authorization json for Oauth
type GoogleJSON struct {
	Web struct {
		ClientID string `json:"client_id"`
		// ProjectID               string `json:"project_id"`
		// AuthURI                 string `json:"auth_uri"`
		// TokenURI                string `json:"token_uri"`
		// AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
		ClientSecret string `json:"client_secret"`
	} `json:"web"`
}

// ReadJSON reads and returns a struct with the client ID and client secret read from our oauth json
func ReadJSON() GoogleJSON {
	// we may or may not want to make this not a file, but this works for now
	jsonFile, err := ioutil.ReadFile("./auth.json")
	if err != nil {
		log.Print(err)
		log.Fatal("Cannot read auth.json")
	}
	var parsed GoogleJSON
	err = json.Unmarshal(jsonFile, &parsed)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}
