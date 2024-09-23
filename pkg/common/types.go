package common

import "github.com/go-resty/resty/v2"

type MethodType string

const (
	PostMethod   MethodType = "POST"
	GetMethod    MethodType = "GET"
	PutDelete    MethodType = "PUT"
	DeleteMethod MethodType = "DELETE"
)

type SDKConfig struct {
	Client *resty.Client
	ApiKey string
}

type Error struct {
	Error string
}

type HelperMeta struct {
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
}

type AccountUpdateRequest struct {
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
}

type PublicDataAsset struct {
	CreatedBy      string        `json:"created_by"`
	ExpirationDate string        `json:"expiration_date"`
	Fid            string        `json:"fid"`
	Size           int           `json:"size"`
	Type           string        `json:"type"`
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	UpdatedAt      string        `json:"updated_at"`
	Acl            []interface{} `json:"acl"`
	CreatedAt      string        `json:"created_at"`
	DataModelId    int           `json:"data_model_id"`
	Tags           []interface{} `json:"tags"`
	TransactionId  string        `json:"transaction_id"`
}

type ResponsesMessageResponse struct {
	Message string `json:"message"`
}

type HelperPaginatedResponse[T any] struct {
	Meta  HelperMeta  `json:"meta"`
	Data  T           `json:"data"`
	Links HelperLinks `json:"links"`
}

type ACLRequest struct {
	Address string        `json:"address"`
	Roles   []interface{} `json:"roles"`
}

type AccessLevel string

const (
	RoleView   AccessLevel = "view"
	RoleUpdate AccessLevel = "update"
	RoleDelete AccessLevel = "delete"
	RoleShare  AccessLevel = "share"
)

type CreateDataAssetRequest struct {
	Acl            []interface{}          `json:"acl"`
	Claim          map[string]interface{} `json:"claim"`
	DataModelId    int                    `json:"data_model_id"`
	ExpirationDate string                 `json:"expiration_date"`
	Name           string                 `json:"name"`
	Tags           []interface{}          `json:"tags"`
}

type PublicACL struct {
	DataAssetId   int           `json:"data_asset_id"`
	Roles         []interface{} `json:"roles"`
	SolanaAddress string        `json:"solana_address"`
	UpdatedAt     string        `json:"updated_at"`
	Address       string        `json:"address"`
	CreatedAt     string        `json:"created_at"`
}

type MyAccountResponse struct {
	Did               string `json:"did"`
	ProfilePicture    string `json:"profile_picture"`
	StorageSize       int    `json:"storage_size"`
	UpdatedAt         string `json:"updated_at"`
	Username          string `json:"username"`
	UsernameUpdatedAt string `json:"username_updated_at"`
	WalletAddress     string `json:"wallet_address"`
	CreatedAt         string `json:"created_at"`
}

type ShareDataAssetRequest struct {
	Addresses []interface{} `json:"addresses"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UpdateDataAssetRequest struct {
	Claim          map[string]interface{} `json:"claim"`
	ExpirationDate string                 `json:"expiration_date"`
	Name           string                 `json:"name"`
}

type DataModelRequest struct {
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        []interface{}          `json:"tags"`
	Title       string                 `json:"title"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type HelperLinks struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type AccountCreateRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	Username      string `json:"username"`
	WalletAddress string `json:"wallet_address"`
}

type AuthRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type DataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type DataModel struct {
	Id          int                    `json:"id"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        []interface{}          `json:"tags"`
	UpdatedAt   string                 `json:"updated_at"`
	CreatedAt   string                 `json:"created_at"`
	CreatedBy   string                 `json:"created_by"`
	DeletedAt   string                 `json:"deleted_at"`
	Description string                 `json:"description"`
	Title       string                 `json:"title"`
}
