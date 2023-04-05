package dto

import (
	"github.com/gin-gonic/gin"
)

type ServiceDTO interface {
	ParseRequest(c *gin.Context)
	Validate()
}
