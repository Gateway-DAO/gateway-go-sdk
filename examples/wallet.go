package examples

import (
	"fmt"
	"log"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg"
)

func ExampleAddWallet(sdk *pkg.SDK) {

	address := "0xYourEthereumAddress"

	myAccount, err := sdk.Account.Wallet.Add(address)
	if err != nil {
		log.Fatalf("Failed to add wallet: %v", err)
	}

	fmt.Println("Added Wallet Account:", myAccount)
}

func ExampleRemoveWallet(sdk *pkg.SDK) {

	address := "0xYourEthereumAddress"

	myAccount, err := sdk.Account.Wallet.Remove(address)
	if err != nil {
		log.Fatalf("Failed to remove wallet: %v", err)
	}

	fmt.Println("Removed Wallet Account:", myAccount)
}

func RunWallet() {
	sdk := pkg.NewSDK(pkg.SDKConfig{WalletDetails: pkg.WalletDetails{PrivateKey: "", WalletType: services.Ethereum}})

	ExampleAddWallet(sdk)
	ExampleRemoveWallet(sdk)
}
