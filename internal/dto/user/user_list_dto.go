package user_dto

import (
	"myapiproject/internal/dto"
	"net/url"
	"strconv"
)

type UserListDTO struct {
	dto.ServiceDTO
	Sort  string
	Page  int
	Limit int
}

func (dto *UserListDTO) ParseRequest(values url.Values) {
	dto.Sort = values.Get("sort")

	page, err := strconv.Atoi(values.Get("page"))
	if err == nil {
		dto.Page = page
	}

	limit, err := strconv.Atoi(values.Get("limit"))

	if err == nil {
		dto.Limit = limit
	}
}

func (dto *UserListDTO) Validate() {

}
