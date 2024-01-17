package models

type Tenant struct {
	DefaultBaseModel
	Name        string `json:"name" gorm:"column:name;not null" validate:"required"`
	Description string `json:"description" gorm:"column:description" validate:""`
	PrimaryHost string `json:"primary_host" gorm:"column:primary_host;not null" validate:"required"`
	Hosts       string `json:"hosts" gorm:"column:hosts;not null" validate:"required"` // separated by comma
	OwnerID     uint   `json:"owner_id" gorm:"column:owner_id" validate:"required"`
}

const TENANT_TABLE_NAME = "tenants"

func (*Tenant) TableName() string {
	return TENANT_TABLE_NAME
}
