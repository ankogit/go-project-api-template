package v1

import (
	"github.com/gin-gonic/gin"
	user_dto "myapiproject/internal/dto/user"
	"myapiproject/internal/transport/rest/middleware"
	"myapiproject/internal/transport/rest/response"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {

	users := api.Group("/", middleware.AuthUser(h.services.TokenManager))
	{
		auth := users.Group("/")
		{
			auth.GET("/users", h.getUsers)
		}
	}

}

func (h *Handler) getUsers(c *gin.Context) {
	//var response ResponseData

	err := c.Request.ParseForm()
	if err != nil {
		response.JsonResponse(c.Writer, response.ResponseData{
			Code:         http.StatusBadRequest,
			ClientErrors: err,
		})
		return
	}

	ulDto := user_dto.UserListDTO{}
	ulDto.ParseRequest(c.Request.Form)

	users, err := h.services.Users.GetList(ulDto)
	if err != nil {
		response.JsonResponse(c.Writer, response.ResponseData{
			Code:         http.StatusNotFound,
			ClientErrors: err,
		})
		return
	}

	response.JsonResponse(c.Writer, response.ResponseData{
		Data: users,
		Meta: struct {
			Total       interface{} `json:"total"`
			CurrentPage int         `json:"current_page"`
		}{
			CurrentPage: ulDto.Page,
		},
	})
}
