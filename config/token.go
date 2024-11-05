package config

import (
	"Expire/model"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type TokenDetails struct {
	Token      *string
	Identifier string
	Email      string
	UserID     string
	ExpiresIn  *int64
}

func CreateToken(user *model.User, ttl time.Duration, privateKey string) (*TokenDetails, error) {
	now := time.Now().UTC()
	td := &TokenDetails{
		ExpiresIn: new(int64),
		Token:     new(string),
	}
	*td.ExpiresIn = now.Add(ttl).Unix()
	td.Identifier = uuid.NewV4().String()
	td.UserID = user.ID.String()
	td.Email = user.Email

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token private key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("create: parse token private key: %w", err)
	}

	atClaims := &jwt.MapClaims{
		"data": map[string]string{
			"email": user.Email,
			"id":    td.Identifier,
		},
		"exp": td.ExpiresIn,
		"iat": now.Unix(),
		"iss": "expire.com",
	}
	*td.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("create: sign token: %w", err)
	}
	return td, nil
}
