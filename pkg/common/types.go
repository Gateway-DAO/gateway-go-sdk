package common

import "github.com/go-resty/resty/v2"

type MethodType string
type WalletSignMessageType struct {
	Signature  []byte
	SigningKey string
}

type SDKConfig struct {
	Client *resty.Client
	ApiKey string
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

type HelperMeta struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

type AuthRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type CreateDataAssetRequest struct {
	Acl            []ACLRequest           `json:"acl"`
	Claim          map[string]interface{} `json:"claim"`
	DataModelId    int                    `json:"data_model_id"`
	ExpirationDate string                 `json:"expiration_date"`
	Name           string                 `json:"name"`
	Tags           []string               `json:"tags"`
}

type MyAccountResponse struct {
	ProfilePicture    string          `json:"profile_picture"`
	StorageSize       int             `json:"storage_size"`
	UpdatedAt         string          `json:"updated_at"`
	Username          string          `json:"username"`
	UsernameUpdatedAt string          `json:"username_updated_at"`
	WalletAddresses   []WalletAddress `json:"wallet_addresses"`
	CreatedAt         string          `json:"created_at"`
	Did               string          `json:"did"`
}

type UpdateDataAssetRequest struct {
	Name           string                 `json:"name"`
	Claim          map[string]interface{} `json:"claim"`
	ExpirationDate string                 `json:"expiration_date"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type PublicDataAsset struct {
	ExpirationDate string      `json:"expiration_date"`
	Size           int         `json:"size"`
	TransactionId  string      `json:"transaction_id"`
	Acl            []PublicACL `json:"acl"`
	CreatedBy      string      `json:"created_by"`
	Fid            string      `json:"fid"`
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	Tags           []string    `json:"tags"`
	Type           string      `json:"type"`
	UpdatedAt      string      `json:"updated_at"`
	CreatedAt      string      `json:"created_at"`
	DataModelId    int         `json:"data_model_id"`
}

type ShareDataAssetRequest struct {
	Addresses []string `json:"addresses"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type AccessLevel string

const (
	RoleView   AccessLevel = "view"
	RoleUpdate AccessLevel = "update"
	RoleDelete AccessLevel = "delete"
	RoleShare  AccessLevel = "share"
)

type AccountCreateRequest struct {
	Signature     string `json:"signature"`
	Username      string `json:"username"`
	WalletAddress string `json:"wallet_address"`
	Message       string `json:"message"`
}

type DataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type DeleteACLRequest struct {
	Addresses []string `json:"addresses"`
}

type PublicAccountResponse struct {
	Did            string `json:"did"`
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
}

type HelperPaginatedResponse[T any] struct {
	Meta  HelperMeta  `json:"meta"`
	Data  T           `json:"data"`
	Links HelperLinks `json:"links"`
}

type ACLRequest struct {
	Roles   []AccessLevel `json:"roles"`
	Address string        `json:"address"`
}

type AccountUpdateRequest struct {
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
}

type DataModel struct {
	CreatedBy   string                 `json:"created_by"`
	Id          int                    `json:"id"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        []string               `json:"tags"`
	Title       string                 `json:"title"`
	UpdatedAt   string                 `json:"updated_at"`
	DeletedAt   string                 `json:"deleted_at"`
}

type DataModelRequest struct {
	Tags        []string               `json:"tags"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
}

type PublicACL struct {
	CreatedAt     string   `json:"created_at"`
	DataAssetId   int      `json:"data_asset_id"`
	Did           string   `json:"did"`
	IsAuthority   bool     `json:"is_authority"`
	Roles         []string `json:"roles"`
	SolanaAddress string   `json:"solana_address"`
	UpdatedAt     string   `json:"updated_at"`
	Address       string   `json:"address"`
}

type WalletAddress struct {
	Id        int    `json:"id"`
	UpdatedAt string `json:"updated_at"`
	AccountId int    `json:"account_id"`
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	CreatedAt string `json:"created_at"`
}

type WalletCreateRequest struct {
	Address string `json:"address"`
}

type ResponsesMessageResponse struct {
	Message string `json:"message"`
}
