package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	AuthorizationHeader = "Authorization"

	UserCtx = "userId"
)

func GetUserIdAuthorization(c *gin.Context) (uint, error) {
	if userId, exists := c.Get(UserCtx); exists {
		userId := userId.(string)

		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			return 0, err
		}
		return uint(userIdInt), nil
	}
	return 0, errors.New("unauthorized")
}
