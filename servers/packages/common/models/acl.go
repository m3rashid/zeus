package models

type AclRelation string

const (
	ACL_RELATION_OWNER   AclRelation = "owner"
	ACL_RELATION_MANAGER AclRelation = "manager"
	ACL_RELATION_EDITOR  AclRelation = "editor"
	ACL_RELATION_VIEWER  AclRelation = "viewer"
)

type Acl struct {
	IDBaseModel
	Entity   string      `json:"property" gorm:"column:property;not null" validate:"required"`
	Relation AclRelation `json:"relation" gorm:"column:relation;not null" validate:"required"`
	UserSet  string      `json:"user_set" gorm:"column:user_set;not null" validate:"required"`
}

const ACL_TABLE_NAME = "acls"

func (*Acl) TableName() string {
	return ACL_TABLE_NAME
}
