package models

type User struct {
	DefaultBaseModel
	Name           string `json:"name" gorm:"column:name;not null" validate:"required"`
	Email          string `json:"email" gorm:"column:email;unique;not null" validate:"required"`
	Password       string `json:"password" gorm:"column:password;not null" validate:"required"`
	OTP            string `json:"otp" gorm:"column:otp;not null" validate:"required"`
	UserVerified   bool   `json:"user_verified" gorm:"column:user_verified;default:false" validate:""`
	OrganizationID uint   `json:"organization_id" gorm:"column:organization_id" validate:""` // if present, it means user belongs to an organization
}

const USER_TABLE_NAME = "users"

func (*User) TableName() string {
	return USER_TABLE_NAME
}
