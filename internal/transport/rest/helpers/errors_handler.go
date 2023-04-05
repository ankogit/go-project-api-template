package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapiproject/internal/transport/rest/response"
	"net/http"
)

func ErrorsHandle(c *gin.Context, err error) {
	// TODO: добавить обработку вариантов ошибок
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.JsonResponse(c.Writer, response.ResponseData{
			Code: http.StatusNotFound,
		})
	} else {
		response.JsonResponse(c.Writer, response.ResponseData{
			Code: http.StatusInternalServerError,
			Text: err.Error(),
		})
	}

}
