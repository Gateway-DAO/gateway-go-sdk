package client

import "github.com/go-resty/resty/v2"

type WalletSignMessageType struct {
	Signature  string
	SigningKey string
}

type Config struct {
	Client *resty.Client
}

type Error struct {
	Error string
}

type HelperLinks struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type AcceptedDataAssetResponse struct {
	AcceptedBy  string `json:"accepted_by"`
	DataAssetId int    `json:"data_asset_id"`
}

type AccountUpdateRequest struct {
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
}

type CreateDataAssetRequest struct {
	Acl            []ACLRequest           `json:"acl"`
	Claim          map[string]interface{} `json:"claim"`
	DataModelId    int                    `json:"data_model_id"`
	ExpirationDate string                 `json:"expiration_date"`
	Name           string                 `json:"name"`
	Tags           []string               `json:"tags"`
}

type PublicDataAsset struct {
	Acl            []PublicACL `json:"acl"`
	ExpirationDate string      `json:"expiration_date"`
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	Tags           []string    `json:"tags"`
	UpdatedAt      string      `json:"updated_at"`
	CreatedAt      string      `json:"created_at"`
	CreatedBy      string      `json:"created_by"`
	DataModelId    int         `json:"data_model_id"`
	Fid            string      `json:"fid"`
	Size           int         `json:"size"`
	TransactionId  string      `json:"transaction_id"`
	Type           string      `json:"type"`
}

type WalletCreateRequest struct {
	Address string `json:"address"`
}

type DataModelCreateRequest struct {
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        []string               `json:"tags"`
	Title       string                 `json:"title"`
}

type PublicACL struct {
	DataAssetId   int      `json:"data_asset_id"`
	Did           string   `json:"did"`
	IsAuthority   bool     `json:"is_authority"`
	Roles         []string `json:"roles"`
	SolanaAddress string   `json:"solana_address"`
	UpdatedAt     string   `json:"updated_at"`
	Address       string   `json:"address"`
	CreatedAt     string   `json:"created_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ComputeRequestResponse struct {
	AcceptedDataAssets []AcceptedDataAssetResponse `json:"accepted_data_assets"`
	ComputeFieldName   string                      `json:"compute_field_name"`
	CreatedAt          string                      `json:"created_at"`
	DataModelId        int                         `json:"data_model_id"`
	UpdatedAt          string                      `json:"updated_at"`
	ComputeOperation   string                      `json:"compute_operation"`
	CreatedBy          string                      `json:"created_by"`
	Description        string                      `json:"description"`
	Id                 int                         `json:"id"`
	Title              string                      `json:"title"`
}

type DeleteACLRequest struct {
	Addresses []string `json:"addresses"`
}

type DataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type ResponsesMessageResponse struct {
	Message string `json:"message"`
}

type TypesAccessLevel string

const (
	RoleView   TypesAccessLevel = "view"
	RoleUpdate TypesAccessLevel = "update"
	RoleDelete TypesAccessLevel = "delete"
	RoleShare  TypesAccessLevel = "share"
)

type DataModelUpdateRequest struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Id          int                    `json:"id"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        []string               `json:"tags"`
}

type ShareDataAssetRequest struct {
	Addresses []string `json:"addresses"`
}

type HelperPaginatedResponse[T any] struct {
	Data  T           `json:"data"`
	Links HelperLinks `json:"links"`
	Meta  HelperMeta  `json:"meta"`
}

type ModelWalletAddress struct {
	AccountId int    `json:"account_id"`
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	UpdatedAt string `json:"updated_at"`
}

type PublicAccountResponse struct {
	Username       string `json:"username"`
	Did            string `json:"did"`
	ProfilePicture string `json:"profile_picture"`
}

type UpdateDataAssetRequest struct {
	Claim          map[string]interface{} `json:"claim"`
	ExpirationDate string                 `json:"expiration_date"`
	Name           string                 `json:"name"`
}

type AccountCreateRequest struct {
	WalletAddress string `json:"wallet_address"`
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	Username      string `json:"username"`
}

type DataModelResponse struct {
	Id          int                    `json:"id"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        []string               `json:"tags"`
	Title       string                 `json:"title"`
	UpdatedAt   string                 `json:"updated_at"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type MyAccountResponse struct {
	Did               string               `json:"did"`
	ProfilePicture    string               `json:"profile_picture"`
	StorageSize       int                  `json:"storage_size"`
	UpdatedAt         string               `json:"updated_at"`
	Username          string               `json:"username"`
	UsernameUpdatedAt string               `json:"username_updated_at"`
	WalletAddresses   []ModelWalletAddress `json:"wallet_addresses"`
	CreatedAt         string               `json:"created_at"`
}

type TypesComputeOperation string

const (
	ComputeOperationAdd                TypesComputeOperation = "add"
	ComputeOperationSubtract           TypesComputeOperation = "subtract"
	ComputeOperationMultiply           TypesComputeOperation = "multiply"
	ComputeOperationDivide             TypesComputeOperation = "divide"
	ComputeOperationSum                TypesComputeOperation = "sum"
	ComputeOperationGreaterThan        TypesComputeOperation = "greater_than"
	ComputeOperationGreaterThanOrEqual TypesComputeOperation = "greater_than_or_equal"
	ComputeOperationLessThan           TypesComputeOperation = "less_than"
	ComputeOperationEqual              TypesComputeOperation = "equal"
	ComputeOperationNotEqual           TypesComputeOperation = "not_equal"
)

type ACLRequest struct {
	Address string             `json:"address"`
	Roles   []TypesAccessLevel `json:"roles"`
}

type AuthRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type ComputeRequestCreateRequest struct {
	ComputeFieldName string      `json:"compute_field_name"`
	ComputeOperation interface{} `json:"compute_operation"`
	DataModelId      int         `json:"data_model_id"`
	Description      string      `json:"description"`
	Title            string      `json:"title"`
}

type HelperMeta struct {
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
}
