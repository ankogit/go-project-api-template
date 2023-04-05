package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myapiproject/internal/models"
	"myapiproject/internal/service"
	"myapiproject/internal/transport/rest/helpers"
	"myapiproject/internal/transport/rest/response"
	nHttp "net/http"
	"strconv"
	"time"
)

type refreshInput struct {
	RefreshToken string `json:"token" form:"token" binding:"required"`
}

type loginInput struct {
	UserId string `json:"user-id" form:"user-id" binding:"required"`
}

type tokenResponse struct {
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
	ExpiresIn    time.Duration `json:"expires_in"`
}

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/refresh", h.refresh)
	}
}

func (h *Handler) login(c *gin.Context) {
	var input loginInput
	if _, err := helpers.BindData(c, &input); err != nil {
		helpers.ErrorsHandle(c, err)
		return
	}

	userId, err := strconv.Atoi(input.UserId)
	if err != nil {
		response.NewResponse(c, nHttp.StatusBadRequest, "Uncorrected user-id")
		return
	}

	resp, err := h.services.Users.Login(service.LoginInput{
		UserId: uint(userId),
	})

	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			response.NewResponse(c, nHttp.StatusBadRequest, err.Error())

			return
		}

		if errors.Is(err, models.ErrStudentBlocked) {
			response.NewResponse(c, nHttp.StatusForbidden, err.Error())

			return
		}

		response.NewResponse(c, nHttp.StatusInternalServerError, err.Error())

		return
	}

	response.JsonResponse(c.Writer, response.ResponseData{
		Code: nHttp.StatusOK,
		Data: tokenResponse{
			AccessToken:  resp.AccessToken,
			RefreshToken: resp.RefreshToken,
		},
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var input refreshInput
	if _, err := helpers.BindData(c, &input); err != nil {
		helpers.ErrorsHandle(c, err)
		return
	}

	resp, err := h.services.Users.RefreshToken(service.RefreshInput{
		Token: input.RefreshToken,
	})

	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			response.NewResponse(c, nHttp.StatusBadRequest, err.Error())

			return
		}

		if errors.Is(err, models.ErrStudentBlocked) {
			response.NewResponse(c, nHttp.StatusForbidden, err.Error())

			return
		}

		response.NewResponse(c, nHttp.StatusInternalServerError, err.Error())

		return
	}

	//response.JsonResponse(c, tokenResponse{
	//	AccessToken:  resp.AccessToken,
	//	RefreshToken: resp.RefreshToken,
	//}, nHttp.StatusOK)

	response.JsonResponse(c.Writer, response.ResponseData{
		Code: nHttp.StatusOK,
		Data: tokenResponse{
			AccessToken:  resp.AccessToken,
			RefreshToken: resp.RefreshToken,
			ExpiresIn:    15 * time.Minute,
		},
	})
}
