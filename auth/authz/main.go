package authz

import (
	"auth/models"
	"fmt"

	"gorm.io/gorm"
)

func HandleAuthorization(
	userID uint,
	resourceID uint,
	resourceType string,
	requiredAcl models.AclRelation,
	db *gorm.DB,
) bool {
	entity := fmt.Sprintf("%s:%d", resourceType, resourceID)
	userSet := fmt.Sprintf("user:%d", userID)

	sql := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE (
			(entity=? AND relation>=? AND user_set LIKE ?) OR
			(entity=? AND relation LIKE ? AND user_set LIKE ?) OR
			(entity=? AND relation LIKE ? AND user_set LIKE ?)
		) OR (entity LIKE ? AND relation>=? AND user_set LIKE ?)`, models.ACL_TABLE_NAME)

	var count uint
	err := db.Exec(
		sql,
		entity,
		requiredAcl,
		userSet,
		entity,
		"%:"+fmt.Sprint(requiredAcl)+":%",
		"%"+userSet+"%",
		entity,
		"%:"+fmt.Sprint(requiredAcl)+":%",
		"%"+userSet+"%",
		"%:"+entity+"%",
		requiredAcl, "%"+userSet+"%",
	).Find(&count).Error
	if err != nil {
		return false
	}

	// debug this
	return count > 0

	/*
		1. Search by entity to get a list of acls
		2. loop over the entities if the actual resolves to true
		3. Loop over the Nested sets
	*/
}
