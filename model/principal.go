package model

type Principal struct {
	EmailVerified bool    `json:"email_verified"`
	AtHash        string  `json:"at_hash"`
	Locale        string  `json:"locale"`
	Exp           float64 `json:"exp"`
	Azp           string  `json:"azp"`
	Aud           string  `json:"aud"`
	Email         string  `json:"email"`
	Iss           string  `json:"iss"`
	Name          string  `json:"name"`
	GivenName     string  `json:"given_name"`
	Iat           float64 `json:"iat"`
	Sub           string  `json:"sub"`
	Picture       string  `json:"picture"`
	FamilyName    string  `json:"family_name"`
	Jti           string  `json:"jti"`
}