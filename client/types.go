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

type WalletCreateRequest struct {
	Address string `json:"address"`
}

type HelperPaginatedResponse[T any] struct {
	Data  T           `json:"data"`
	Links HelperLinks `json:"links"`
	Meta  HelperMeta  `json:"meta"`
}

type ResponsesMessageResponse struct {
	Message string `json:"message"`
}

type ComputeRequestAcceptRequest struct {
	DataAssetId int `json:"data_asset_id"`
}

type ComputeRequestResponse struct {
	Description           *string                      `json:"description"`
	Title                 *string                      `json:"title"`
	ComputeFieldName      *string                      `json:"compute_field_name"`
	CreatedAt             *string                      `json:"created_at"`
	ComputeOperationParam *int                         `json:"compute_operation_param"`
	CreatedBy             *string                      `json:"created_by"`
	DataModelId           *int                         `json:"data_model_id"`
	Id                    *int                         `json:"id"`
	UpdatedAt             *string                      `json:"updated_at"`
	AcceptedDataAssets    *[]AcceptedDataAssetResponse `json:"accepted_data_assets"`
	ComputeOperation      *string                      `json:"compute_operation"`
}

type AuthRequest struct {
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type PublicAccountResponse struct {
	Did            string  `json:"did"`
	ProfilePicture *string `json:"profile_picture"`
	Username       string  `json:"username"`
}

type PublicACL struct {
	IsAuthority   *bool    `json:"is_authority"`
	Roles         []string `json:"roles"`
	SolanaAddress string   `json:"solana_address"`
	UpdatedAt     *string  `json:"updated_at"`
	Address       string   `json:"address"`
	CreatedAt     *string  `json:"created_at"`
	Did           *string  `json:"did"`
}

type HelperLinks struct {
	Last     string `json:"last"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	First    string `json:"first"`
}

type ModelWalletAddress struct {
	AccountId int     `json:"account_id"`
	Address   string  `json:"address"`
	Chain     string  `json:"chain"`
	CreatedAt string  `json:"created_at"`
	Id        int     `json:"id"`
	UpdatedAt *string `json:"updated_at"`
}

type ComputeRequestReceivedResponse struct {
	ComputeOperationParam *int    `json:"compute_operation_param"`
	CreatedAt             *string `json:"created_at"`
	Description           *string `json:"description"`
	Title                 *string `json:"title"`
	ComputeFieldName      *string `json:"compute_field_name"`
	CreatedBy             *string `json:"created_by"`
	DataAssetsIds         *[]int  `json:"data_assets_ids"`
	DataModelId           *int    `json:"data_model_id"`
	Id                    *int    `json:"id"`
	UpdatedAt             *string `json:"updated_at"`
	ComputeOperation      *string `json:"compute_operation"`
}

type MyAccountResponse struct {
	Username          string               `json:"username"`
	UsernameUpdatedAt string               `json:"username_updated_at"`
	WalletAddresses   []ModelWalletAddress `json:"wallet_addresses"`
	CreatedAt         string               `json:"created_at"`
	Did               string               `json:"did"`
	ProfilePicture    *string              `json:"profile_picture"`
	StorageSize       int                  `json:"storage_size"`
	UpdatedAt         string               `json:"updated_at"`
}

type DataModelCreateRequest struct {
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
	Tags        *[]string              `json:"tags"`
	Title       string                 `json:"title"`
}

type ShareDataAssetRequest struct {
	Addresses []string `json:"addresses"`
}

type ACLRequest struct {
	Address string             `json:"address"`
	Roles   []TypesAccessLevel `json:"roles"`
}

type ComputingProcessResponse struct {
	ComputeRequest *int    `json:"compute_request"`
	ComputeResult  *string `json:"compute_result"`
	ComputeStatus  *string `json:"compute_status"`
	CreatedBy      *string `json:"created_by"`
	Id             *int    `json:"id"`
}

type DataAssetIDRequestAndResponse struct {
	Id int `json:"id"`
}

type DeleteACLRequest struct {
	Addresses []string `json:"addresses"`
}

type PublicDataAsset struct {
	Type           string      `json:"type"`
	CreatedAt      *string     `json:"created_at"`
	CreatedBy      string      `json:"created_by"`
	DataModelId    *int        `json:"data_model_id"`
	ExpirationDate *string     `json:"expiration_date"`
	Fid            string      `json:"fid"`
	Name           string      `json:"name"`
	TransactionId  string      `json:"transaction_id"`
	UpdatedAt      *string     `json:"updated_at"`
	Acl            []PublicACL `json:"acl"`
	Id             int         `json:"id"`
	Size           int         `json:"size"`
	Tags           *[]string   `json:"tags"`
}

type AccountCreateRequest struct {
	Signature     string `json:"signature"`
	Username      string `json:"username"`
	WalletAddress string `json:"wallet_address"`
	Message       string `json:"message"`
}

type CreateDataAssetRequest struct {
	Acl            *[]ACLRequest           `json:"acl"`
	Claim          *map[string]interface{} `json:"claim"`
	DataModelId    *int                    `json:"data_model_id"`
	ExpirationDate *string                 `json:"expiration_date"`
	Name           string                  `json:"name"`
	Tags           *[]string               `json:"tags"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UpdateDataAssetRequest struct {
	Claim          *map[string]interface{} `json:"claim"`
	ExpirationDate *string                 `json:"expiration_date"`
	Name           *string                 `json:"name"`
}

type HelperMeta struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

type TypesAccessLevel string

const (
	RoleView   TypesAccessLevel = "view"
	RoleUpdate TypesAccessLevel = "update"
	RoleDelete TypesAccessLevel = "delete"
	RoleShare  TypesAccessLevel = "share"
)

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
	ComputeOperationExponential        TypesComputeOperation = "exponential"
	ComputeOperationRemainder          TypesComputeOperation = "remainder"
)

type AccountUpdateRequest struct {
	ProfilePicture *string `json:"profile_picture"`
	Username       *string `json:"username"`
}

type DataModelResponse struct {
	Schema      map[string]interface{} `json:"schema"`
	Tags        *[]string              `json:"tags"`
	Title       string                 `json:"title"`
	UpdatedAt   string                 `json:"updated_at"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
	Id          int                    `json:"id"`
}

type DataModelUpdateRequest struct {
	Tags        *[]string              `json:"tags"`
	Title       *string                `json:"title"`
	Description *string                `json:"description"`
	Id          int                    `json:"id"`
	Schema      map[string]interface{} `json:"schema"`
}

type AcceptedDataAssetResponse struct {
	AcceptedBy  *string `json:"accepted_by"`
	DataAssetId *int    `json:"data_asset_id"`
}

type ComputeRequestCreateRequest struct {
	Description           string      `json:"description"`
	Title                 string      `json:"title"`
	ComputeFieldName      string      `json:"compute_field_name"`
	ComputeOperation      interface{} `json:"compute_operation"`
	ComputeOperationParam *int        `json:"compute_operation_param"`
	DataModelId           int         `json:"data_model_id"`
}
