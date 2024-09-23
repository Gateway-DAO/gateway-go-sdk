package datamodels

import "github.com/Gateway-DAO/gateway-go-sdk/pkg/common"

type DataModel interface {
	GetDataModels(page int, page_size int) (common.HelperPaginatedResponse[[]common.DataModel], error)
	CreateDataModel(dataModelInput common.DataModelRequest) (common.DataModel, error)
	GetMyDataModels(page int, page_size int) (common.HelperPaginatedResponse[[]common.DataModel], error)
	GetDataModelById(id int64) (common.DataModel, error)
	UpdateDataModelById(id int64) (common.DataModel, error)
}

type DataModelImpl struct {
	Config common.SDKConfig
}

func NewDataModelImpl(config common.SDKConfig) *DataModelImpl {
	return &DataModelImpl{
		Config: config,
	}
}

