package dataassets

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"time"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type FileResponse struct {
	FileName    string
	FileContent []byte
	FileType    string
}

type DataAsset interface {
	Upload(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error)
	UploadFile(fileName string, fileContent []byte, aclList *[]common.ACLRequest, expirationDate *time.Time) (common.DataAssetIDRequestAndResponse, error)
	GetCreatedByMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error)
	GetReceivedByMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error)
	Get(id int64) (common.PublicDataAsset, error)
	UpdateAsset(id string, dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error)
	UpdateFile(id string, fileName string, fileContent []byte, aclList *[]common.ACLRequest, expirationDate *time.Time) (common.PublicDataAsset, error)
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

func (u *DataAssetImpl) Get(id int64) (common.PublicDataAsset, error) {
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

func (u *DataAssetImpl) GetReceivedByMe(page int, page_size int) (common.HelperPaginatedResponse[[]common.PublicDataAsset], error) {
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

func (u *DataAssetImpl) Upload(dataAssetInput common.CreateDataAssetRequest) (common.DataAssetIDRequestAndResponse, error) {
	var id common.DataAssetIDRequestAndResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&dataAssetInput).SetResult(&id).SetError(&error).Post(common.CreateANewDataAsset)

	if err != nil {
		return id, err
	}

	if res.IsError() {
		return id, errors.New(error.Error)
	}

	return id, nil
}

func toRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func (u *DataAssetImpl) UploadFile(fileName string, fileContent []byte, aclList *[]common.ACLRequest, expirationDate *time.Time) (common.DataAssetIDRequestAndResponse, error) {
	var id common.DataAssetIDRequestAndResponse
	var error common.Error

	formData := make(map[string]string)

	if aclList != nil {
		aclJSON, err := json.Marshal(aclList)
		if err != nil {
			return id, err
		}
		formData["acl"] = string(aclJSON)
	}

	if expirationDate != nil {
		formData["expiration_date"] = toRFC3339(*expirationDate)
	}

	req := u.Config.Client.R().SetFileReader("data", fileName, bytes.NewReader(fileContent))

	if len(formData) > 0 {
		req = req.SetFormData(formData)
	}

	res, err := req.SetResult(&id).SetError(&error).Post(common.CreateANewDataAsset)

	if err != nil {
		return id, err
	}

	if res.IsError() {
		return id, errors.New(error.Error)
	}

	return id, nil
}

func (u *DataAssetImpl) UpdateAsset(id string, dataAssetInput common.UpdateDataAssetRequest) (common.PublicDataAsset, error) {
	var asset common.PublicDataAsset
	var error common.Error

	res, err := u.Config.Client.R().SetPathParam("id", id).SetBody(&dataAssetInput).SetResult(&asset).SetError(&error).Put(common.UpdateDataAssetByID)

	if err != nil {
		return asset, err
	}

	if res.IsError() {
		return asset, errors.New(error.Error)
	}

	return asset, nil
}

func (u *DataAssetImpl) UpdateFile(id string, fileName string, fileContent []byte, aclList *[]common.ACLRequest, expirationDate *time.Time) (common.PublicDataAsset, error) {
	var asset common.PublicDataAsset
	var error common.Error

	formData := make(map[string]string)

	if aclList != nil {
		aclJSON, err := json.Marshal(aclList)
		if err != nil {
			return asset, err
		}
		formData["acl"] = string(aclJSON)
	}

	if expirationDate != nil {
		formData["expiration_date"] = toRFC3339(*expirationDate)
	}

	req := u.Config.Client.R().SetFileReader("data", fileName, bytes.NewReader(fileContent))

	if len(formData) > 0 {
		req = req.SetFormData(formData)
	}

	res, err := req.SetPathParam("id", id).SetResult(&asset).SetError(&error).Put(common.UpdateDataAssetByID)

	if err != nil {
		return asset, err
	}

	if res.IsError() {
		return asset, errors.New(error.Error)
	}

	return asset, nil

}

func (u *DataAssetImpl) DeleteAsset(id int64) (common.MessageResponse, error) {
	var message common.MessageResponse
	var error common.Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%v", id)).SetResult(&message).SetError(&error).Delete(common.DeleteDataAssetByID)

	if err != nil {
		return message, err
	}

	if res.IsError() {
		return message, errors.New(error.Error)
	}

	return message, nil
}

func (u *DataAssetImpl) Share(id int64, shareDetails []common.ShareDataAssetRequest) ([]common.PublicACL, error) {

	var acl []common.PublicACL
	var error common.Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%v", id)).SetBody(&shareDetails).SetResult(&acl).SetError(&error).Post(common.ShareDataAssetByID)
	
	if err != nil {
		return acl, err
	}

	if res.IsError() {
		return acl, errors.New(error.Error)
	}

	return acl, nil
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
