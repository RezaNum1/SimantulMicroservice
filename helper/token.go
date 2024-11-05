package helper

import (
	"Expire/config"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// import (
// 	"Expire/model"
// 	"context"
// 	"encoding/base64"
// 	"fmt"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	uuid "github.com/satori/go.uuid"
// )

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ValidateTokenFormat(token string, publicKey string) (*jwt.Token, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	extractedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	return extractedToken, err
}

// Utilize puiblic key to verify JWT and Extract Corresponding Payload
func ExtractIdentifierFromToken(token string, publicKey string) (*string, error) { // Return Identifier / Id inside data object from access / refresh token

	parsedToken, err := ValidateTokenFormat(token, publicKey)

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	fmt.Println("Claims:", claims)

	data := claims["data"]
	dataModel := data.(map[string]interface{})
	idValue := dataModel["id"].(string)
	return &idValue, nil
}

func GetUserIdFromIdentifier(identifier string) (string, error) {
	contextTodo := context.TODO()
	userId, err := config.RedisClient.Get(contextTodo, identifier).Result()

	return userId, err
}

func GetAccessTokenPublicKey() string {
	env, _ := config.LoadConfig(".")
	return env.AccessTokenPublicKey
}

func GetRefreshTokenPublicKey() string {
	env, _ := config.LoadConfig(".")
	return env.RefreshTokenPublicKey
}

func GetUserId(ctx *gin.Context) (*string, error) {
	token := ExtractToken(ctx)
	identifier, errId := ExtractIdentifierFromToken(token, GetAccessTokenPublicKey())
	if errId != nil {
		return nil, errId
	}
	userId, errGetUserId := GetUserIdFromIdentifier(*identifier)
	if errGetUserId != nil {
		return nil, errId
	}
	return &userId, nil
}
