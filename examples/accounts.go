package examples

import (
	"fmt"
	"log"

	
)

func ExampleLogin(sdk *client.SDK) {

	message := "example-message"
	signature := "example-signature"
	walletAddress := "0xYourEthereumAddress"

	token, err := sdk.Auth.Login(message, signature, walletAddress)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	fmt.Println("JWT Token:", token)
}

func ExampleGetMessage(sdk *client.SDK) {

	message, err := sdk.Auth.GetMessage()
	if err != nil {
		log.Fatalf("Failed to get message: %v", err)
	}

	fmt.Println("Sign Message:", message)
}

func ExampleGetRefreshToken(sdk *client.SDK) {

	refreshToken, err := sdk.Auth.GetRefreshToken()
	if err != nil {
		log.Fatalf("Failed to get refresh token: %v", err)
	}

	fmt.Println("Refresh Token:", refreshToken)
}

func RunAccounts() {
	sdk := client.NewSDK(client.SDKConfig{WalletDetails: client.WalletDetails{PrivateKey: "", WalletType: services.Ethereum}})

	ExampleLogin(sdk)
	ExampleGetMessage(sdk)
	ExampleGetRefreshToken(sdk)
}
