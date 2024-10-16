package dataassets

import (
	"errors"

	"github.com/Gateway-DAO/gateway-go-sdk/pkg/common"
)

type ACL interface {
	Add(id int64, aclList []common.ACLRequest) (common.PublicACL, error)
	Update(id int64, aclList []common.ACLRequest) (common.PublicACL, error)
	Delete(id int64, aclList []common.ACLRequest) (string, error)
}

type ACLImpl struct {
	Config common.SDKConfig
}

func NewACLImpl(config common.SDKConfig) *ACLImpl {
	return &ACLImpl{
		Config: config,
	}
}

func (u *ACLImpl) Add(id int64, aclList []common.ACLRequest) (common.PublicACL, error) {
	var publicACL common.PublicACL
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&aclList).SetResult(&publicACL).SetError(&error).Post(common.AssignACLItemsToDataAsset)

	if err != nil {
		return publicACL, err
	}

	if res.IsError() {
		return publicACL, errors.New(error.Error)
	}

	return publicACL, nil

}

func (u *ACLImpl) Update(id int64, aclList []common.ACLRequest) (common.PublicACL, error) {
	var publicACL common.PublicACL
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&aclList).SetResult(&publicACL).SetError(&error).Put(common.UpdateACLItemsToDataAsset)

	if err != nil {
		return publicACL, err
	}

	if res.IsError() {
		return publicACL, errors.New(error.Error)
	}

	return publicACL, nil
}

func (u *ACLImpl) Delete(id int64, aclList []common.ACLRequest) (string, error) {
	var response common.MessageResponse
	var error common.Error

	res, err := u.Config.Client.R().SetBody(&aclList).SetResult(&response).SetError(&error).Delete(common.DeleteAssignedRoleByACL)

	if err != nil {
		return response.Message, err
	}

	if res.IsError() {
		return response.Message, errors.New(error.Error)
	}

	return response.Message, nil
}
