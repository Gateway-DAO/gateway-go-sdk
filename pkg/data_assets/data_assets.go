package dataassets

import (
	"net/http"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type DataAssets struct {
	client *http.Client
	apiKey string
}

func New(client *http.Client, apiKey string) *DataAssets {
	return &DataAssets{
		client: client,
		apiKey: apiKey,
	}
}

func (m *DataAssets) GetData(id string) (string, error) {

	return services.MakeRequest(m.client, "GET", common.GetDataAssetByID, id)
}
