package client

import (
	"errors"
	"fmt"
)

type DataModel interface {
	GetAll(page int, page_size int) (HelperPaginatedResponse[[]DataModelResponse], error)
	Create(dataModelInput DataModelCreateRequest) (DataModelResponse, error)
	GetMy(page int, page_size int) (HelperPaginatedResponse[[]DataModelResponse], error)
	GetById(id int64) (DataModelResponse, error)
	Update(id int64, dataModelInput DataModelUpdateRequest) (DataModelResponse, error)
}

type DataModelImpl struct {
	Config Config
}

func NewDataModelImpl(config Config) *DataModelImpl {
	return &DataModelImpl{
		Config: config,
	}
}

func (u *DataModelImpl) GetAll(page int, page_size int) (HelperPaginatedResponse[[]DataModelResponse], error) {

	var dataModels HelperPaginatedResponse[[]DataModelResponse]
	var error Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&dataModels).SetError(&error).Get(GetDataModels)

	if err != nil {
		return dataModels, err
	}

	if res.IsError() {
		return dataModels, errors.New(error.Error)
	}

	return dataModels, nil
}

func (u *DataModelImpl) GetMy(page int, page_size int) (HelperPaginatedResponse[[]DataModelResponse], error) {
	var dataModels HelperPaginatedResponse[[]DataModelResponse]
	var error Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&dataModels).SetError(&error).Get(GetDataModelsByUser)

	if err != nil {
		return dataModels, err
	}

	if res.IsError() {
		return dataModels, errors.New(error.Error)
	}

	return dataModels, nil
}

func (u *DataModelImpl) GetById(id int64) (DataModelResponse, error) {

	var dataModel DataModelResponse
	var error Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%d", id)).SetResult(&dataModel).SetError(&error).Get(GetDataModelByID)

	if err != nil {
		return dataModel, err
	}

	if res.IsError() {
		return dataModel, errors.New(error.Error)
	}

	return dataModel, nil

}

func (u *DataModelImpl) Create(dataModelInput DataModelCreateRequest) (DataModelResponse, error) {
	var dataModelCreated DataModelResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&dataModelInput).SetResult(&dataModelCreated).SetError(&error).Post(CreateDataModel)

	if err != nil {
		return dataModelCreated, err
	}

	if res.IsError() {
		return dataModelCreated, errors.New(error.Error)
	}

	return dataModelCreated, nil
}

func (u *DataModelImpl) Update(id int64, dataModelInput DataModelUpdateRequest) (DataModelResponse, error) {
	var dataModelUpdated DataModelResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&dataModelInput).SetResult(&dataModelUpdated).SetError(&error).Put(UpdateDataModel)

	if err != nil {
		return dataModelUpdated, err
	}

	if res.IsError() {
		return dataModelUpdated, errors.New(error.Error)
	}

	return dataModelUpdated, nil

}
