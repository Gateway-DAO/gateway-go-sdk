package dataassets

import (
	"net/http"

	"github.com/Gateway-DAO/gateway-go-sdk/internal/services"
	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type DataAssets struct {
	Client *http.Client
	ApiKey string
}

func (m *DataAssets) GetData(id string) (string, error) {

	return services.MakeRequest(m.Client, "GET", common.GetDataAssetByID, id)
}
