package dataassets

import (
	"errors"
	"fmt"
	"mime"
	"os"
	"path/filepath"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type FileResponse struct {
	FileName    string
	FileContent []byte
	FileType    string
}

type DataAsset interface {
	CreateStructured(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error)
	CreateNonStructured(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error)
	GetCreatedByMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error)
	GetReceivedToMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error)
	GetDetail(id int64) (common.PublicDataAsset, error)
	UpdateStructured(dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error)
	UpdateNonStructured(dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error)
	DeleteAsset(id int64) (common.MessageResponse, error)
	Download(id int64) (*FileResponse, error)
	Share(id int64, shareDetails []common.ShareDataAssetRequest) ([]common.PublicACL, error)
}

type DataAssetImpl struct {
	Config common.SDKConfig
}

func NewDataAssetImpl(config common.SDKConfig) *DataAssetImpl {
	return &DataAssetImpl{
		Config: config,
	}
}

func (u *DataAssetImpl) GetDetail(id int64) (common.PublicDataAsset, error) {
	var asset common.PublicDataAsset
	var error common.Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%d", id)).SetResult(&asset).SetError(&error).Get(common.GetDataAssetByID)

	if err != nil {
		return asset, err
	}

	if res.IsError() {
		return asset, errors.New(error.Error)
	}

	return asset, nil
}

func (u *DataAssetImpl) GetCreatedByMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error) {

	var assets common.HelperPaginatedResponse[[]common.PublicDataAsset]
	var error common.Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&assets).SetError(&error).Get(common.GetCreatedDataAssets)

	if err != nil {
		return assets, err
	}

	if res.IsError() {
		return assets, errors.New(error.Error)
	}

	return assets, nil
}

func (u *DataAssetImpl) GetReceivedToMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error) {
	var assets common.HelperPaginatedResponse[[]common.PublicDataAsset]
	var error common.Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&assets).SetError(&error).Get(common.GetReceivedDataAssets)

	if err != nil {
		return assets, err
	}

	if res.IsError() {
		return assets, errors.New(error.Error)
	}

	return assets, nil
}

func (u *DataAssetImpl) CreateStructured(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error) {

}

func (u *DataAssetImpl) CreateNonStructured(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error) {

}

func (u *DataAssetImpl) UpdateStructured(dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error) {

}

func (u *DataAssetImpl) UpdateNonStructured(dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error) {

}

func (u *DataAssetImpl) DeleteAsset(id int64) (common.MessageResponse, error) {

}

func (u *DataAssetImpl) Share(id int64, shareDetails []common.ShareDataAssetRequest) ([]common.PublicACL, error) {

}

func (u *DataAssetImpl) Download(id int64) (*FileResponse, error) {
	resp, err := u.Config.Client.R().
		SetOutput("temporary-file").
		Get(common.DownloadDataAssetByID)

	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}

	contentDisposition := resp.Header().Get("Content-Disposition")
	var fileName string
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err == nil {
			fileName = params["filename"]
		}
	}

	if fileName == "" {
		fileName = filepath.Base(resp.Request.URL)
	}

	fileContent, err := os.ReadFile("temporary-file")
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	defer os.Remove("temporary-file")

	fileType := filepath.Ext(fileName)

	return &FileResponse{
		FileName:    fileName,
		FileContent: fileContent,
		FileType:    fileType,
	}, nil

}
