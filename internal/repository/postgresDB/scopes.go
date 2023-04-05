package postgresDB

import (
	"gorm.io/gorm"
	"myapiproject/pkg/pagination"
)

func paginate(value interface{}, pagination *pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(value).Count(&totalRows)
	pagination.Total = totalRows

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
