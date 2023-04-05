package response

import (
	"github.com/alioygur/gores"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code         int         `json:"code,omitempty" default:"200"`
	Text         string      `json:"text,omitempty"`
	ClientErrors interface{} `json:"errors,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Meta         interface{} `json:"meta,omitempty"`
	Links        interface{} `json:"links,omitempty"`
	Error        error       `json:"-"`
}

type response struct {
	Message string `json:"message"`
}

func NewResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{message})
}

func JsonResponse(w http.ResponseWriter, response ResponseData) {
	if response.Code == http.StatusInternalServerError {
		// log error
	}

	gores.JSON(w, response.Code, response)
}
