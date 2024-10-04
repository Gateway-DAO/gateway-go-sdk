package common

const (
	GetCreatedDataAssets = "/data-assets/created"
	DeleteAssignedRoleByACL = "/data-assets/{id}/acl/delete"
	GetDataModelByID = "/data-models/{id}"
	UpdateDataModel = "/data-models/{id}"
	CreateANewDataAsset = "/data-assets"
	AddWallet = "/accounts/me/wallets"
	RemoveWallet = "/accounts/me/wallets/{address}"
	AuthenticateAccount = "/auth"
	GetDataAssetByID = "/data-assets/{id}"
	UpdateDataAssetByID = "/data-assets/{id}"
	DeleteDataAssetByID = "/data-assets/{id}"
	CreateAccount = "/accounts"
	ShareDataAssetByID = "/data-assets/{id}/share"
	GetDataModels = "/data-models"
	CreateDataModel = "/data-models"
	AssignACLItemsToDataAsset = "/data-assets/{id}/acl"
	UpdateACLItemsToDataAsset = "/data-assets/{id}/acl"
	GenerateSignMessage = "/auth/message"
	RefreshToken = "/auth/refresh-token"
	GetReceivedDataAssets = "/data-assets/received"
	DownloadDataAssetByID = "/data-assets/{id}/download"
	GetDataModelsByUser = "/data-models/me"
	GetMyAccount = "/accounts/me"
	UpdateAccount = "/accounts/me"
)

