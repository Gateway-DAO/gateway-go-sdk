package helpers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/auth"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckJWTTokenExpiration(tokenString string) (bool, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("token is not valid")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return false, nil
	}

	return true, nil
}

func IssueJWT(client resty.Client, wallet services.Wallet) (string, error) {
	auth := auth.NewAuthImpl(common.SDKConfig{Client: &client, ApiKey: ""})

	message, messageErr := auth.GetMessage()
	if messageErr != nil {
		return "", messageErr
	}
	log.Println("in message", message)

	signatureDetails, signingErr := wallet.SignMessage(message)
	if signingErr != nil {
		return "", signingErr
	}

	jwt, authErr := auth.Login(message, string(signatureDetails.Signature), signatureDetails.SigningKey)
	if authErr != nil {
		return "", authErr
	}
	return jwt, nil
}

var UNPROTECTED_ROUTES = []string{common.GenerateSignMessage,
	common.RefreshToken, common.AuthenticateAccount}

func AuthMiddleware(params services.MiddlewareParams) resty.RequestMiddleware {
	return func(c *resty.Client, r *resty.Request) error {
		log.Println("here", r.URL, UNPROTECTED_ROUTES[2])

		for _, route := range UNPROTECTED_ROUTES {
			if route == r.URL {
				return nil
			}
		}
		accessToken := r.Header.Get("Authorization")
		log.Println("existing token", accessToken)
		if accessToken == "" {
			log.Println("empty access token")

			newToken, err := IssueJWT(*params.Client, &params.Wallet)
			if err != nil {
				return fmt.Errorf("failed to issue new token: %v", err)
			}
			accessToken = newToken
		} else {
			isValid, _ := CheckJWTTokenExpiration(accessToken)
			log.Println("in jwt check", isValid)

			if !isValid {
				newToken, err := IssueJWT(*params.Client, &params.Wallet)
				if err != nil {
					return fmt.Errorf("failed to issue new token: %v", err)
				}
				accessToken = newToken
			}

		}
		r.Header.Set("Authorization", accessToken)

		return nil
	}
}
