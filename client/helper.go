package client

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckJWTTokenExpiration(tokenString string) (bool, error) {
	claims := &jwt.RegisteredClaims{}

	_, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(tokenString, claims)

	if err != nil {
		return false, err
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return false, nil
	}

	return true, nil
}

func IssueJWT(client resty.Client, wallet Wallet) (string, error) {
	auth := NewAuthImpl(Config{Client: &client})

	message, messageErr := auth.GetMessage()
	if messageErr != nil {
		return "", messageErr
	}

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

var UNPROTECTED_ROUTES = []string{GenerateSignMessage,
	RefreshToken, AuthenticateAccount}

func AuthMiddleware(params MiddlewareParams) resty.RequestMiddleware {
	return func(c *resty.Client, r *resty.Request) error {
		for _, route := range UNPROTECTED_ROUTES {
			if route == r.URL {
				return nil
			}
		}
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			newToken, err := IssueJWT(*params.Client, &params.Wallet)
			if err != nil {
				return fmt.Errorf("failed to issue new token: %v", err)
			}
			accessToken = newToken
		} else {
			isValid, _ := CheckJWTTokenExpiration(accessToken)

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
