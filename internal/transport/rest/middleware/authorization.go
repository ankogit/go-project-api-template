package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myapiproject/internal/transport/rest/helpers"
	"myapiproject/internal/transport/rest/response"
	"myapiproject/pkg/auth"
	"net/http"
	"strings"
)

func AuthUser(tokenManager auth.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseAuthHeader(c, tokenManager)
		if err != nil {
			response.NewResponse(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		c.Set(helpers.UserCtx, id)
	}
}

func OptionalAuthUser(tokenManager auth.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := parseAuthHeader(c, tokenManager)
		if id != "" {
			c.Set(helpers.UserCtx, id)
		}
	}
}

func parseAuthHeader(c *gin.Context, tokenManager auth.TokenManager) (string, error) {
	header := c.GetHeader(helpers.AuthorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}
	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return tokenManager.ParseAccessToken(headerParts[1])
}
