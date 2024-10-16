package helpers

import (
	"errors"
	"time"

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

var UNPROTECTED_ROUTES = []string{common.GenerateSignMessage,
	common.RefreshToken, common.AuthenticateAccount}

func AuthMiddleware(c *resty.Client, r *resty.Request) error {
	for _, route := range UNPROTECTED_ROUTES {
		if route == r.URL {
			return nil
		}
	}
	accessToken := r.Header.Get("Authorization")

	if accessToken == "" {
		// newToken, err := IssueJWT("username")
		newToken := ""
		// if err != nil {
		// 	return fmt.Errorf("failed to issue new token: %v", err)
		// }
		accessToken = newToken
	} else {
		isValid, _ := CheckJWTTokenExpiration(accessToken)
		if !isValid {

		}

	}
	r.Header.Set("Authorization", accessToken)

	return nil
}
