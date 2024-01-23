package models

import "common/server"

const CLIENT_TABLE_NAME = "clients"

type Client struct {
	server.DefaultBaseModel
	ClientID           string       `json:"client_id" gorm:"column:client_id;not null;unique" validate:"required"`
	ClientSecret       string       `json:"client_secret" gorm:"column:client_secret;unique;not null" validate:"required"`
	SuccessRedirectUri string       `json:"success_redirect_uri" gorm:"column:success_redirect_uri;not null" validate:"required"`
	FailureRedirectUri string       `json:"failure_redirect_uri" gorm:"column:failure_redirect_uri;not null" validate:"required"`
	AppName            string       `json:"client_app_name" gorm:"column:client_app_name;not null" validate:"required"`
	AppLogoUrl         string       `json:"client_app_logo_url" gorm:"column:client_app_logo_url" validate:""`
	CreatedByUserID    uint         `json:"createdby_user_id" gorm:"column:createdby_user_id;not null" validate:"required"`
	Scopes             server.JSONB `json:"scopes" gorm:"column:scopes;not null" validate:"required"`
	// structure: { "contacts": []string{"read", "create"} }
}

func (*Client) TableName() string {
	return CLIENT_TABLE_NAME
}

const CONNECTED_CLIENT_TABLE_NAME = "connected_apps"

type ConnectedClients struct {
	ClientID uint `json:"client_id" gorm:"column:client_id;primary_key;not null" validate:"required"` // id of the client (not the client_id)
	UserID   uint `json:"user_id" gorm:"column:user_id;primary_key;not null" validate:"required"`
}

func (*ConnectedClients) TableName() string {
	return CONNECTED_CLIENT_TABLE_NAME
}
