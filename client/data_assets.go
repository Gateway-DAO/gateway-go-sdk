package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"time"
)

type FileResponse struct {
	FileName    string
	FileContent []byte
	FileType    string
}

type DataAsset interface {
	Upload(dataAssetInput CreateDataAssetRequest) (DataAssetIDRequestAndResponse, error)
	UploadFile(fileName string, fileContent []byte, aclList *[]ACLRequest, expirationDate *time.Time) (DataAssetIDRequestAndResponse, error)
	GetCreatedByMe(page int, page_size int) (HelperPaginatedResponse[[]PublicDataAsset], error)
	GetReceivedByMe(page int, page_size int) (HelperPaginatedResponse[[]PublicDataAsset], error)
	Get(id int64) (PublicDataAsset, error)
	UpdateAsset(id string, dataAssetInput UpdateDataAssetRequest) (PublicDataAsset, error)
	UpdateFile(id string, fileName string, fileContent []byte, aclList *[]ACLRequest, expirationDate *time.Time) (PublicDataAsset, error)
	DeleteAsset(id int64) (MessageResponse, error)
	Download(id int64) (*FileResponse, error)
	Share(id int64, shareDetails []ShareDataAssetRequest) ([]PublicACL, error)
}

type DataAssetImpl struct {
	Config Config
}

func NewDataAssetImpl(config Config) *DataAssetImpl {
	return &DataAssetImpl{
		Config: config,
	}
}

func (u *DataAssetImpl) Get(id int64) (PublicDataAsset, error) {
	var asset PublicDataAsset
	var error Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%d", id)).SetResult(&asset).SetError(&error).Get(GetDataAssetByID)

	if err != nil {
		return asset, err
	}

	if res.IsError() {
		return asset, errors.New(error.Error)
	}

	return asset, nil
}

func (u *DataAssetImpl) GetCreatedByMe(page int, page_size int) (HelperPaginatedResponse[[]PublicDataAsset], error) {

	var assets HelperPaginatedResponse[[]PublicDataAsset]
	var error Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&assets).SetError(&error).Get(GetCreatedDataAssets)

	if err != nil {
		return assets, err
	}

	if res.IsError() {
		return assets, errors.New(error.Error)
	}

	return assets, nil
}

func (u *DataAssetImpl) GetReceivedByMe(page int, page_size int) (HelperPaginatedResponse[[]PublicDataAsset], error) {
	var assets HelperPaginatedResponse[[]PublicDataAsset]
	var error Error

	res, err := u.Config.Client.R().SetQueryParams(map[string]string{
		"page":      fmt.Sprintf("%d", page),
		"page_size": fmt.Sprintf("%d", page_size),
	}).SetResult(&assets).SetError(&error).Get(GetReceivedDataAssets)

	if err != nil {
		return assets, err
	}

	if res.IsError() {
		return assets, errors.New(error.Error)
	}

	return assets, nil
}

func (u *DataAssetImpl) Upload(dataAssetInput CreateDataAssetRequest) (DataAssetIDRequestAndResponse, error) {
	var id DataAssetIDRequestAndResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&dataAssetInput).SetResult(&id).SetError(&error).Post(CreateANewDataAsset)

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

func (u *DataAssetImpl) UploadFile(fileName string, fileContent []byte, aclList *[]ACLRequest, expirationDate *time.Time) (DataAssetIDRequestAndResponse, error) {
	var id DataAssetIDRequestAndResponse
	var error Error

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

	res, err := req.SetResult(&id).SetError(&error).Post(CreateANewDataAsset)

	if err != nil {
		return id, err
	}

	if res.IsError() {
		return id, errors.New(error.Error)
	}

	return id, nil
}

func (u *DataAssetImpl) UpdateAsset(id string, dataAssetInput UpdateDataAssetRequest) (PublicDataAsset, error) {
	var asset PublicDataAsset
	var error Error

	res, err := u.Config.Client.R().SetPathParam("id", id).SetBody(&dataAssetInput).SetResult(&asset).SetError(&error).Put(UpdateDataAssetByID)

	if err != nil {
		return asset, err
	}

	if res.IsError() {
		return asset, errors.New(error.Error)
	}

	return asset, nil
}

func (u *DataAssetImpl) UpdateFile(id string, fileName string, fileContent []byte, aclList *[]ACLRequest, expirationDate *time.Time) (PublicDataAsset, error) {
	var asset PublicDataAsset
	var error Error

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

	res, err := req.SetPathParam("id", id).SetResult(&asset).SetError(&error).Put(UpdateDataAssetByID)

	if err != nil {
		return asset, err
	}

	if res.IsError() {
		return asset, errors.New(error.Error)
	}

	return asset, nil

}

func (u *DataAssetImpl) DeleteAsset(id int64) (MessageResponse, error) {
	var message MessageResponse
	var error Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%v", id)).SetResult(&message).SetError(&error).Delete(DeleteDataAssetByID)

	if err != nil {
		return message, err
	}

	if res.IsError() {
		return message, errors.New(error.Error)
	}

	return message, nil
}

func (u *DataAssetImpl) Share(id int64, shareDetails []ShareDataAssetRequest) ([]PublicACL, error) {

	var acl []PublicACL
	var error Error

	res, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%v", id)).SetBody(&shareDetails).SetResult(&acl).SetError(&error).Post(ShareDataAssetByID)

	if err != nil {
		return acl, err
	}

	if res.IsError() {
		return acl, errors.New(error.Error)
	}

	return acl, nil
}

func (u *DataAssetImpl) Download(id int64) (*FileResponse, error) {
	resp, err := u.Config.Client.R().SetPathParam("id", fmt.Sprintf("%v", id)).
		SetOutput("temporary-file").
		Get(DownloadDataAssetByID)

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
