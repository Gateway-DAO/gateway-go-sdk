package examples

import (
	"fmt"
	"log"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg"
)

func ExampleLogin(sdk *pkg.SDK) {
	message := "example-message"
	signature := "example-signature"
	walletAddress := "0xYourEthereumAddress"

	token, err := sdk.Auth.Login(message, signature, walletAddress)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	fmt.Println("JWT Token:", token)
}

func ExampleGetMessage(sdk *pkg.SDK) {

	message, err := sdk.Auth.GetMessage()
	if err != nil {
		log.Fatalf("Failed to get message: %v", err)
	}

	fmt.Println("Sign Message:", message)
}

func ExampleGetRefreshToken(sdk *pkg.SDK) {

	refreshToken, err := sdk.Auth.GetRefreshToken()
	if err != nil {
		log.Fatalf("Failed to get refresh token: %v", err)
	}

	fmt.Println("Refresh Token:", refreshToken)
}

func RunAccounts() {
	sdk := pkg.NewSDK(pkg.SDKConfig{WalletDetails: pkg.WalletDetails{PrivateKey: "", WalletType: services.Ethereum}})

	ExampleLogin(sdk)
	ExampleGetMessage(sdk)
	ExampleGetRefreshToken(sdk)
}
