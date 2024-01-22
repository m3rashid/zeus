package flow

import (
	"common/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const localUsersCookieName = "local_users"

func getLocalUserIDsFromCookie(ctx *fiber.Ctx) []uint {
	var users []uint
	userIDsString := ctx.Cookies(localUsersCookieName, "")
	userIDsStringArray := strings.Split(userIDsString, ",")
	for _, userIDString := range userIDsStringArray {
		u64, err := strconv.ParseUint(userIDString, 10, 32)
		if err != nil {
			return users
		}
		users = append(users, uint(u64))
	}

	return users
}

func setLocalUsersCookie(ctx *fiber.Ctx, userIDs []uint) {
	var userIDStringArray []string
	for _, userId := range userIDs {
		userIDStringArray = append(userIDStringArray, strconv.FormatUint(uint64(userId), 10))
	}

	// TODO: also hash this
	userIDsString := strings.Join(utils.RemoveDuplicates(userIDStringArray), ",")
	ctx.Cookie(&fiber.Cookie{
		HTTPOnly: true,
		Name:     localUsersCookieName,
		Value:    userIDsString,
		Domain:   "localhost", // TODO: handle this for deployments
	})
}

func addUserIDToCookie(ctx *fiber.Ctx, userID uint) {
	existingUserIDs := getLocalUserIDsFromCookie(ctx)
	existingUserIDs = append(existingUserIDs, userID)
	setLocalUsersCookie(ctx, existingUserIDs)
}

func removeUserIDFromCookie(ctx *fiber.Ctx, userID uint) {
	var newUserIDs []uint
	existingUserIDs := getLocalUserIDsFromCookie(ctx)
	for _, existingUserID := range existingUserIDs {
		if existingUserID != userID {
			newUserIDs = append(newUserIDs, existingUserID)
		}
	}

	setLocalUsersCookie(ctx, newUserIDs)
}
