package datamodels

import (
	"errors"
	"fmt"

	"gateway-go-sdk/pkg/common"
)

type DataModel interface {
	GetAll(page int, page_size int) (common.HelperPaginatedResponse[[]common.DataModel], error)
	Create(dataModelInput common.DataModelRequest) (common.DataModel, error)
	GetMy(page int, page_size int) (common.HelperPaginatedResponse[[]common.DataModel], error)
	GetById(id int64) (common.DataModel, error)
	Update(id int64, dataModelInput common.DataModelRequest) (common.DataModel, error)
}

type DataModelImpl struct {
	Config common.SDKConfig
}

func NewDataModelImpl(config common.SDKConfig) *DataModelImpl {
	return &DataModelImpl{
		Config: config,
	}
}

func (u *DataModelImpl) GetAll(page int, page_size int) (common.HelperPaginatedResponse[[]common.DataModel], error) {

	var dataModels common.HelperPaginatedResponse[[]common.DataModel]
	var error common.Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&dataModels).SetError(&error).Get(common.GetDataModels)

	if err != nil {
		return dataModels, err
	}

	if res.IsError() {
		return dataModels, errors.New(error.Error)
	}

	return dataModels, nil
}

func (u *DataModelImpl) GetMy(page int, page_size int) (common.HelperPaginatedResponse[[]common.DataModel], error) {
	var dataModels common.HelperPaginatedResponse[[]common.DataModel]
	var error common.Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&dataModels).SetError(&error).Get(common.GetDataModelsByUser)

	if err != nil {
		return dataModels, err
	}

	if res.IsError() {
		return dataModels, errors.New(error.Error)
	}

	return dataModels, nil
}

func (u *DataModelImpl) GetById(id int64) (common.DataModel, error) {

	var dataModel common.DataModel
	var error common.Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%d", id)).SetResult(&dataModel).SetError(&error).Get(common.GetDataModelByID)

	if err != nil {
		return dataModel, err
	}

	if res.IsError() {
		return dataModel, errors.New(error.Error)
	}

	return dataModel, nil

}

func (u *DataModelImpl) Create(dataModelInput common.DataModelRequest) (common.DataModel, error) {
	var dataModelCreated common.DataModel
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&dataModelInput).SetResult(&dataModelCreated).SetError(&error).Post(common.CreateDataModel)

	if err != nil {
		return dataModelCreated, err
	}

	if res.IsError() {
		return dataModelCreated, errors.New(error.Error)
	}

	return dataModelCreated, nil
}

func (u *DataModelImpl) Update(id int64, dataModelInput common.DataModelRequest) (common.DataModel, error) {
	var dataModelUpdated common.DataModel
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&dataModelInput).SetResult(&dataModelUpdated).SetError(&error).Put(common.UpdateDataModel)

	if err != nil {
		return dataModelUpdated, err
	}

	if res.IsError() {
		return dataModelUpdated, errors.New(error.Error)
	}

	return dataModelUpdated, nil

}
