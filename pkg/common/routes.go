package common

const (
	AddFundsToAccount       = "/accounts/%s/add-funds"
	RefreshToken            = "/auth/refresh-token"
	GetMyDataAssets         = "/data-assets/me"
	DownloadDataAssetByID   = "/data-assets/%s/download"
	CreateANewDataAsset     = "/data-assets"
	GetDataModels           = "/data-models"
	CreateDataModel         = "/data-models"
	GetDataModelByID        = "/data-models/%s"
	UpdateDataModel         = "/data-models/%s"
	CreateAccount           = "/accounts"
	GetMyAccount            = "/accounts/me"
	GenerateSignMessage     = "/auth/message"
	GetDataAssetByID        = "/data-assets/%s"
	UpdateDataAssetByID     = "/data-assets/%s"
	DeleteDataAssetByID     = "/data-assets/%s"
	DeleteAssignedRoleByACL = "/data-assets/%s/delete-assigned-role"
	AuthenticateAccount     = "/auth"
	GetDataModelsByUser     = "/data-models/me"
)
