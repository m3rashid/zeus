package models

type Group struct {
	DefaultBaseModel
	Name        string `json:"name" gorm:"column:name;not null" validate:"required"`
	Description string `json:"description" gorm:"column:description" validate:""`
}

const GROUP_TABLE_NAME = "groups"

func (*Group) TableName() string {
	return GROUP_TABLE_NAME
}

type GroupUserRelationTable struct {
	IDBaseModel
	UserID  uint `json:"user_id" gorm:"column:user_id;not null" validate:"required"`
	GroupID uint `json:"group_id" gorm:"column:group_id;not null" validate:"required"`
}

const GROUP_USER_RELATION_TABLE_NAME = "group_user_relations"

func (*GroupUserRelationTable) TableName() string {
	return GROUP_USER_RELATION_TABLE_NAME
}
