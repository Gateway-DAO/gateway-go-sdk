package client

import (
	"errors"
)

type ACL interface {
	Add(id int64, aclList []ACLRequest) (PublicACL, error)
	Update(id int64, aclList []ACLRequest) (PublicACL, error)
	Delete(id int64, aclList []ACLRequest) (string, error)
}

type ACLImpl struct {
	Config Config
}

func NewACLImpl(config Config) *ACLImpl {
	return &ACLImpl{
		Config: config,
	}
}

func (u *ACLImpl) Add(id int64, aclList []ACLRequest) (PublicACL, error) {
	var publicACL PublicACL
	var error Error

	res, err := u.Config.Client.R().SetBody(&aclList).SetResult(&publicACL).SetError(&error).Post(AssignACLItemsToDataAsset)

	if err != nil {
		return publicACL, err
	}

	if res.IsError() {
		return publicACL, errors.New(error.Error)
	}

	return publicACL, nil

}

func (u *ACLImpl) Update(id int64, aclList []ACLRequest) (PublicACL, error) {
	var publicACL PublicACL
	var error Error

	res, err := u.Config.Client.R().SetBody(&aclList).SetResult(&publicACL).SetError(&error).Put(UpdateACLItemsToDataAsset)

	if err != nil {
		return publicACL, err
	}

	if res.IsError() {
		return publicACL, errors.New(error.Error)
	}

	return publicACL, nil
}

func (u *ACLImpl) Delete(id int64, aclList []ACLRequest) (string, error) {
	var response MessageResponse
	var error Error

	res, err := u.Config.Client.R().SetBody(&aclList).SetResult(&response).SetError(&error).Delete(DeleteAssignedRoleByACL)

	if err != nil {
		return response.Message, err
	}

	if res.IsError() {
		return response.Message, errors.New(error.Error)
	}

	return response.Message, nil
}
