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

// check all interfaces with swagger again
type DataAsset interface {
	CreateClaimDataAsset(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error)
	CreateFileDataAsset(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error)
	GetMyDataAssets() (common.HelperPaginatedResponse[[]common.PublicDataAsset], error)
	GetDataAssetByID(id int64) (common.PublicDataAsset, error)
	UpdateClaimDataAsset(dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error)
	UpdateFileDataAsset(dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error)
	DeleteDataAssetById(id int64) (common.MessageResponse, error)
	AssignACLToDataAsset(id int64, aclList []common.ACLRequest) (common.PublicACL, error)
	UpdateACLToDataAsset(id int64, aclList []common.ACLRequest) (common.PublicACL, error)
	DownloadDataAssetById(id int64) (*FileResponse, error)
	ShareDataAssetById(id int64, shareDetails []common.ShareDataAssetRequest) ([]common.PublicACL, error)
}

type DataAssetImpl struct {
	Config common.SDKConfig
}

func NewDataAssetImpl(config common.SDKConfig) *DataAssetImpl {
	return &DataAssetImpl{
		Config: config,
	}
}

// check for not initiaized values
func (u *DataAssetImpl) GetDataAssetByID(id int64) (common.PublicDataAsset, error) {
	var asset common.PublicDataAsset
	var error common.Error

	res, err := u.Config.Client.R().SetPathParams(map[string]string{
		"id": fmt.Sprintf("%d", id),
	}).SetResult(&asset).SetError(&error).Get(common.GetDataAssetByID)

	if err != nil {
		return common.PublicDataAsset{}, err
	}

	if res.IsError() {
		return common.PublicDataAsset{}, errors.New(error.Error)
	}

	return asset, nil
}

func (u *DataAssetImpl) DownloadDataAssetById(id int64) (*FileResponse, error) {
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
