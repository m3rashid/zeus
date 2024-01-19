package models

type AclRelation int
type AclRelationString string

var ACL_RELATIONS = []AclRelationString{"owner", "manager", "editor", "viewer"}

const (
	ACL_RELATION_OWNER   AclRelation = 1 // owner
	ACL_RELATION_MANAGER AclRelation = 2 // manager
	ACL_RELATION_EDITOR  AclRelation = 3 // editor
	ACL_RELATION_VIEWER  AclRelation = 4 // viewer
)

type Acl struct {
	IDBaseModel
	Entity   string      `json:"entity" gorm:"column:entity;not null" validate:"required"`
	Relation AclRelation `json:"relation" gorm:"column:relation;not null" validate:"required"`
	UserSet  string      `json:"user_set" gorm:"column:user_set;not null" validate:"required"`
}

const ACL_TABLE_NAME = "acls"

func (*Acl) TableName() string {
	return ACL_TABLE_NAME
}

func (a AclRelation) ToValue() AclRelationString {
	return ACL_RELATIONS[a-1]
}

func (a AclRelationString) ToInt() int {
	switch a {
	case "owner":
		return int(ACL_RELATION_OWNER)
	case "manager":
		return int(ACL_RELATION_MANAGER)
	case "editor":
		return int(ACL_RELATION_EDITOR)
	case "viewer":
		return int(ACL_RELATION_VIEWER)
	default:
		return 0
	}
}
