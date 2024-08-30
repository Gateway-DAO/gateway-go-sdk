package main

type ModelDataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type ModelDataModel struct {
	CreatedBy string `json:"created_by"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Id int `json:"id"`
	Schema map[string]interface{} `json:"schema"`
	Tags []interface{} `json:"tags"`
	Title string `json:"title"`
	UpdatedAt string `json:"updated_at"`
}

type ModelMessageResponse struct {
	Message string `json:"message"`
}

type ResponsesEntityRemovedResponse struct {
	Message string `json:"message"`
}

type HelperLinks struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	First string `json:"first"`
	Last string `json:"last"`
}

type ModelAccessLevel string

const (
	Read ModelAccessLevel = "Read"
	Write ModelAccessLevel = "Write"
)

type ModelAccountLedgerCreateResponse struct {
	AccountId int `json:"account_id"`
	Amount float64 `json:"amount"`
	Balance float64 `json:"balance"`
	CreatedAt string `json:"createdAt"`
	Id int `json:"id"`
	TransactionType ModelTransactionType `json:"transaction_type"`
	UpdatedAt string `json:"updatedAt"`
}

type ModelCreateDataAssetRequest struct {
	Tags []interface{} `json:"tags"`
	Acl []interface{} `json:"acl"`
	Claim map[string]interface{} `json:"claim"`
	DataModelId int `json:"data_model_id"`
	ExpirationDate string `json:"expiration_date"`
	Name string `json:"name"`
}

type ModelMyAccountResponse struct {
	CreatedAt string `json:"created_at"`
	Did string `json:"did"`
	ProfilePicture string `json:"profile_picture"`
	UpdatedAt string `json:"updated_at"`
	Username string `json:"username"`
	WalletAddress string `json:"wallet_address"`
}

type ModelPublicRole struct {
	CreatedAt string `json:"created_at"`
	DataAssetId int `json:"data_asset_id"`
	Role string `json:"role"`
	UpdatedAt string `json:"updated_at"`
	WalletAddress string `json:"wallet_address"`
}

type HelperMeta struct {
	ItemsPerPage int `json:"items_per_page"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
}

type HelperPaginatedResponse struct {
	Meta HelperMeta `json:"meta"`
	Data interface{} `json:"data"`
	Links HelperLinks `json:"links"`
}

type ModelAccountLedgerAddRequest struct {
	Amount float64 `json:"amount"`
}

type ModelAuthRequest struct {
	Message string `json:"message"`
	Signature string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type ModelRoleRequest struct {
	Role ModelAccessLevel `json:"role"`
	Wallet string `json:"wallet"`
}

type ModelTokenResponse struct {
	Token string `json:"token"`
}

type ModelTransactionType string

const (
	TransactionTypeDeposit ModelTransactionType = "deposit"
	TransactionTypeWithdrawal ModelTransactionType = "withdrawal"
	TransactionTypePayment ModelTransactionType = "payment"
	TransactionTypeFeePayment ModelTransactionType = "fee_payment"
)

type ModelAccountCreateRequest struct {
	Message string `json:"message"`
	Signature string `json:"signature"`
	Username string `json:"username"`
	WalletAddress string `json:"wallet_address"`
}

type ModelPublicDataAsset struct {
	Roles []interface{} `json:"roles"`
	Size int `json:"size"`
	TransactionId string `json:"transaction_id"`
	UpdatedAt string `json:"updated_at"`
	Issuer ModelMyAccountResponse `json:"issuer"`
	Name string `json:"name"`
	Tags []interface{} `json:"tags"`
	Type string `json:"type"`
	CreatedAt string `json:"created_at"`
	DataModelId int `json:"data_model_id"`
	ExpirationDate string `json:"expiration_date"`
	Fid string `json:"fid"`
}

const (
	GetDataAssetByID = "/data-assets/%s"
	UpdateDataAssetByID = "/data-assets/%s"
	DeleteDataAssetByID = "/data-assets/%s"
	GetDataModels = "/data-models"
	CreateDataModel = "/data-models"
	CreateAccount = "/accounts"
	AuthenticateAccount = "/auth"
	GenerateSignMessage = "/auth/message"
	CreateANewDataAsset = "/data-assets"
	GetMyDataAssets = "/data-assets/me"
	DownloadDataAssetByID = "/data-assets/%s/download"
	GetDataModelsByUser = "/data-models/me"
	GetMyAccount = "/accounts/me"
	UpdateDataModel = "/data-models/%s"
	GetDataModelByID = "/data-models/%s"
	DeleteAssignedRoleByACL = "/data-assets/%s/delete-assigned-role"
	RefreshToken = "/auth/refresh-token"
	AddFundsToAccount = "/accounts/%s/add-funds"
)