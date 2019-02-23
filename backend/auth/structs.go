package auth

// IDTokenClaims is the format for the (useful) info we get back from OIDC.
// This might not work for other providers that aren't google, as some of these fields aren't standard.
type IDTokenClaims struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Nonce         string `json:"nonce"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
}
