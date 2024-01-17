package models

import "time"

type IDBaseModel struct {
	ID uint `gorm:"primary_key;column:id" json:"id"`
}

type DefaultBaseModel struct {
	IDBaseModel
	Deleted   bool      `gorm:"column:deleted;default:false" json:",omitempty" validate:""`
	CreatedAt time.Time `gorm:"column:createdAt; default:current_timestamp" json:"createdAt"`
}
