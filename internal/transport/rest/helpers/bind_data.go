package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"myapiproject/internal/transport/rest/response"
	"net/http"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// BindData is helper function, returns false if data is not bound
func BindData(c *gin.Context, req interface{}) (int, error) {
	err := c.Request.ParseForm()
	if err != nil {
		//response.JsonResponse(c.Writer, response.ResponseData{
		//	Code:         http.StatusBadRequest,
		//	ClientErrors: err,
		//})
		return http.StatusBadRequest, err
	}

	if c.ContentType() != "application/json" && c.ContentType() != "multipart/form-data" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		//response.NewResponse(c, http.StatusBadRequest, msg)
		return http.StatusBadRequest, errors.New(msg)
	}
	// Bind incoming json to struct and check for validation errors
	if err := c.Bind(req); err != nil {
		//log.Printf("Error binding data: %+v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			// could probably extract this, it is also in middleware_auth_user
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			response.NewResponse(c, http.StatusBadRequest, "Invalid request parameters. See invalidArgs")
			return http.StatusBadRequest, errors.New("Invalid request parameters. See invalidArgs")

		}

		// later we'll add code for validating max body size here!

		// if we aren't able to properly extract validation errors,
		// we'll fallback and return an internal server error
		//response.NewResponse(c, http.StatusHTTPVersionNotSupported, err.Error())
		return http.StatusBadRequest, err
	}

	return 0, nil
}
