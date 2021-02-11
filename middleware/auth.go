package middleware

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-jwx/jwk"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"strings"
	"turrium/env"
	"turrium/model"
	"turrium/repository"
)

func parseKeys(path string) map[string]*rsa.PublicKey {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return make(map[string]*rsa.PublicKey)
	}

	set, err := jwk.Parse(data)
	if err != nil {
		return make(map[string]*rsa.PublicKey)
	}

	keymap := make(map[string]*rsa.PublicKey)
	for _, key := range set.Keys {
		if mkey, err := key.Materialize(); err == nil {
			keymap[key.KeyID()] = mkey.(*rsa.PublicKey)
		}
	}

	return keymap
}

func parseAuthorization(header string) string {
	auth := strings.Split(header, " ")
	token := auth[len(auth)-1]
	return token
}

func verifyClaims(claims jwt.MapClaims) bool {
	return claims.VerifyAudience(env.OAUTH_CLIENT_ID, true) && claims.VerifyIssuer(env.OAUTH_CLIENT_ISSUER, true)
}

func extractStringClaim(claims jwt.MapClaims, id string) string {
	value, ok := claims[id].(string)
	if !ok {
		return ""
	}
	return value
}

func extractFloatClaim(claims jwt.MapClaims, id string) float64 {
	value, ok := claims[id].(float64)
	if !ok {
		return 0
	}
	return value
}

func extractBooleanClaim(claims jwt.MapClaims, id string) bool {
	value, ok := claims[id].(bool)
	if !ok {
		return false
	}
	return value
}

func extractPrincipal(claims jwt.MapClaims) model.Principal {
	return model.Principal{
		EmailVerified: extractBooleanClaim(claims, "email_verified"),
		AtHash:        extractStringClaim(claims, "at_hash"),
		Locale:        extractStringClaim(claims, "locale"),
		Exp:           extractFloatClaim(claims, "exp"),
		Azp:           extractStringClaim(claims, "azp"),
		Aud:           extractStringClaim(claims, "aud"),
		Email:         extractStringClaim(claims, "email"),
		Iss:           extractStringClaim(claims, "iss"),
		Name:          extractStringClaim(claims, "name"),
		GivenName:     extractStringClaim(claims, "given_name"),
		Iat:           extractFloatClaim(claims, "iat"),
		Sub:           extractStringClaim(claims, "sub"),
		Picture:       extractStringClaim(claims, "picture"),
		FamilyName:    extractStringClaim(claims, "family_name"),
		Jti:           extractStringClaim(claims, "jti"),
	}
}

func VerifyTokens() gin.HandlerFunc {
	keys := parseKeys("../certs/google.json")
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		authorization := parseAuthorization(header)
		var claims jwt.MapClaims = map[string]interface{}{}
		_, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
			if kid, ok := token.Header["kid"]; ok {
				return keys[kid.(string)], nil
			}
			return nil, errors.New("no matching key found")
		})
		if err != nil || !verifyClaims(claims) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("principal", extractPrincipal(claims))
		c.Next()
	}
}

func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("principal")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		principal, ok := value.(model.Principal)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		allowedUsers := repository.GetUserEmails(bson.M{})
		if _, allowed := allowedUsers[principal.Email]; !allowed {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}