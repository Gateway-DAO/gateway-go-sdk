package common

import "github.com/go-resty/resty/v2"

type MethodType string

type SDKConfig struct {
	Client *resty.Client
	ApiKey string
}

type Error struct {
	Error string
}

type AccountCreateRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	Username      string `json:"username"`
	WalletAddress string `json:"wallet_address"`
}

type DeleteACLRequest struct {
	Addresses []interface{} `json:"addresses"`
}

type PublicDataAsset struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	Fid            string        `json:"fid"`
	Tags           []string      `json:"tags"`
	TransactionId  string        `json:"transaction_id"`
	UpdatedAt      string        `json:"updated_at"`
	DataModelId    int           `json:"data_model_id"`
	ExpirationDate string        `json:"expiration_date"`
	CreatedAt      string        `json:"created_at"`
	Type           string        `json:"type"`
	Size           int           `json:"size"`
	Acl            []interface{} `json:"acl"`
	CreatedBy      string        `json:"created_by"`
}

type DataModelRequest struct {
	Tags        []interface{}          `json:"tags"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
}

type ShareDataAssetRequest struct {
	Addresses []interface{} `json:"addresses"`
}

type UpdateDataAssetRequest struct {
	Claim          map[string]interface{} `json:"claim"`
	ExpirationDate string                 `json:"expiration_date"`
	Name           string                 `json:"name"`
}

type WalletCreateRequest struct {
	Address string `json:"address"`
}

type ResponsesMessageResponse struct {
	Message string `json:"message"`
}

type HelperPaginatedResponse[T any] struct {
	Data  T           `json:"data"`
	Links HelperLinks `json:"links"`
	Meta  HelperMeta  `json:"meta"`
}

type ACLRequest struct {
	Address string        `json:"address"`
	Roles   []interface{} `json:"roles"`
}

type AccountUpdateRequest struct {
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
}

type AuthRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type MyAccountResponse struct {
	WalletAddresses   []interface{} `json:"wallet_addresses"`
	CreatedAt         string        `json:"created_at"`
	Did               string        `json:"did"`
	ProfilePicture    string        `json:"profile_picture"`
	StorageSize       int           `json:"storage_size"`
	UpdatedAt         string        `json:"updated_at"`
	Username          string        `json:"username"`
	UsernameUpdatedAt string        `json:"username_updated_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type WalletAddress struct {
	AccountId int    `json:"account_id"`
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	UpdatedAt string `json:"updated_at"`
}

type HelperLinks struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	First    string `json:"first"`
	Last     string `json:"last"`
}

type HelperMeta struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

type AccessLevel string

const (
	RoleView   AccessLevel = "view"
	RoleUpdate AccessLevel = "update"
	RoleDelete AccessLevel = "delete"
	RoleShare  AccessLevel = "share"
)

type CreateDataAssetRequest struct {
	Name           string                 `json:"name"`
	Tags           []interface{}          `json:"tags"`
	Acl            []interface{}          `json:"acl"`
	Claim          map[string]interface{} `json:"claim"`
	DataModelId    int                    `json:"data_model_id"`
	ExpirationDate string                 `json:"expiration_date"`
}

type DataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type DataModel struct {
	CreatedAt   string                 `json:"created_at"`
	CreatedBy   string                 `json:"created_by"`
	DeletedAt   string                 `json:"deleted_at"`
	Description string                 `json:"description"`
	Tags        []interface{}          `json:"tags"`
	UpdatedAt   string                 `json:"updated_at"`
	Id          int                    `json:"id"`
	Schema      map[string]interface{} `json:"schema"`
	Title       string                 `json:"title"`
}

type PublicACL struct {
	Roles         []interface{} `json:"roles"`
	SolanaAddress string        `json:"solana_address"`
	UpdatedAt     string        `json:"updated_at"`
	Address       string        `json:"address"`
	CreatedAt     string        `json:"created_at"`
	DataAssetId   int           `json:"data_asset_id"`
	IsAuthority   bool          `json:"is_authority"`
}
