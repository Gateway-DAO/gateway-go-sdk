package common

type ModelAccessLevel string

const (
	Read ModelAccessLevel = "Read"
	Write ModelAccessLevel = "Write"
)

type ModelTransactionType string

const (
	TransactionTypeDeposit ModelTransactionType = "deposit"
	TransactionTypeWithdrawal ModelTransactionType = "withdrawal"
	TransactionTypePayment ModelTransactionType = "payment"
	TransactionTypeFeePayment ModelTransactionType = "fee_payment"
)

type HelperMeta struct {
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
}

type ModelAccountLedgerAddRequest struct {
	Amount float64 `json:"amount"`
}

type ModelCreateDataAssetRequest struct {
	Acl []interface{} `json:"acl"`
	Claim map[string]interface{} `json:"claim"`
	DataModelId int `json:"data_model_id"`
	ExpirationDate string `json:"expiration_date"`
	Name string `json:"name"`
	Tags []interface{} `json:"tags"`
}

type ModelMessageResponse struct {
	Message string `json:"message"`
}

type ModelPublicRole struct {
	WalletAddress string `json:"wallet_address"`
	CreatedAt string `json:"created_at"`
	DataAssetId int `json:"data_asset_id"`
	Role string `json:"role"`
	UpdatedAt string `json:"updated_at"`
}

type ResponsesEntityRemovedResponse struct {
	Message string `json:"message"`
}

type HelperLinks struct {
	First string `json:"first"`
	Last string `json:"last"`
	Next string `json:"next"`
	Previous string `json:"previous"`
}

type HelperPaginatedResponse struct {
	Data interface{} `json:"data"`
	Links HelperLinks `json:"links"`
	Meta HelperMeta `json:"meta"`
}

type ModelAccountCreateRequest struct {
	Message string `json:"message"`
	Signature string `json:"signature"`
	Username string `json:"username"`
	WalletAddress string `json:"wallet_address"`
}

type ModelDataModel struct {
	DeletedAt string `json:"deleted_at"`
	Description string `json:"description"`
	Id int `json:"id"`
	Tags []interface{} `json:"tags"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Schema map[string]interface{} `json:"schema"`
	Title string `json:"title"`
}

type ModelRoleRequest struct {
	Role ModelAccessLevel `json:"role"`
	Wallet string `json:"wallet"`
}

type ModelTokenResponse struct {
	Token string `json:"token"`
}

type ModelAccountLedgerCreateResponse struct {
	CreatedAt string `json:"createdAt"`
	Id int `json:"id"`
	TransactionType ModelTransactionType `json:"transaction_type"`
	UpdatedAt string `json:"updatedAt"`
	AccountId int `json:"account_id"`
	Amount float64 `json:"amount"`
	Balance float64 `json:"balance"`
}

type ModelAuthRequest struct {
	Message string `json:"message"`
	Signature string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type ModelDataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type ModelMyAccountResponse struct {
	CreatedAt string `json:"created_at"`
	Did string `json:"did"`
	ProfilePicture string `json:"profile_picture"`
	UpdatedAt string `json:"updated_at"`
	Username string `json:"username"`
	WalletAddress string `json:"wallet_address"`
}

type ModelPublicDataAsset struct {
	ExpirationDate string `json:"expiration_date"`
	Fid string `json:"fid"`
	Issuer ModelMyAccountResponse `json:"issuer"`
	Tags []interface{} `json:"tags"`
	Type string `json:"type"`
	CreatedAt string `json:"created_at"`
	DataModelId int `json:"data_model_id"`
	Name string `json:"name"`
	Roles []interface{} `json:"roles"`
	Size int `json:"size"`
	TransactionId string `json:"transaction_id"`
	UpdatedAt string `json:"updated_at"`
}

