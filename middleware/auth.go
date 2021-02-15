package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-jwx/jwk"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
	"turrium/env"
	"turrium/model"
	"turrium/repository"
	"turrium/structs"
)

func parseKeys() map[string]*rsa.PublicKey {
	set, err := jwk.FetchHTTP(env.JWK_URL)
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
	keys := parseKeys()
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		authorization := parseAuthorization(header)
		var claims jwt.MapClaims = map[string]interface{}{}
		_, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
			if kid, ok := token.Header["kid"]; ok {
				if key, exists := keys[kid.(string)]; exists {
					return key, nil
				}
			}
			return nil, errors.New("token could not be verified")
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, structs.Status{
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Reason: "Unauthorized.",
			})
			return
		}

		if !claims.VerifyAudience(env.OAUTH_CLIENT_ID, true) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, structs.Status{
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Reason: "Invalid claim: aud.",
			})
			return
		}
		if !claims.VerifyIssuer(env.OAUTH_CLIENT_ISSUER, true) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, structs.Status{
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Reason: "Invalid claim: iss.",
			})
		}

		c.Set("principal", extractPrincipal(claims))
		c.Next()
	}
}

func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("principal")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, structs.Status{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Reason: "Nonexistent principal.",
			})
			return
		}

		principal, ok := value.(model.Principal)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, structs.Status{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Reason: "Invalid principal.",
			})
			return
		}

		fmt.Println("======== PRINCIPAL ========")
		fmt.Println("Timestamp: " + time.Now().Format(time.RFC3339))
		JSON, _ := json.Marshal(principal)
		fmt.Println("Principal: " + string(JSON))
		fmt.Println("===========================")

		allowedUsers := repository.GetUserEmails(bson.M{})
		if _, allowed := allowedUsers[principal.Email]; !allowed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, structs.Status{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Reason: "Unauthorized user.",
			})
			return
		}

		c.Next()
	}
}