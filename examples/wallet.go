package examples

import (
	"fmt"
	"log"

	"github.com/Gateway-DAO/gateway-go-sdk/client"
)

func ExampleAddWallet(sdk *client.SDK) {

	address := "0xYourEthereumAddress"

	myAccount, err := sdk.Account.Wallet.Add(address)
	if err != nil {
		log.Fatalf("Failed to add wallet: %v", err)
	}

	fmt.Println("Added Wallet Account:", myAccount)
}

func ExampleRemoveWallet(sdk *client.SDK) {

	address := "0xYourEthereumAddress"

	myAccount, err := sdk.Account.Wallet.Remove(address)
	if err != nil {
		log.Fatalf("Failed to remove wallet: %v", err)
	}

	fmt.Println("Removed Wallet Account:", myAccount)
}

func RunWallet() {
	sdk := client.NewSDK(client.SDKConfig{WalletDetails: client.WalletDetails{PrivateKey: "", WalletType: client.Ethereum}})

	ExampleAddWallet(sdk)
	ExampleRemoveWallet(sdk)
}
