package database

import (
	"math"
	"zup-message-service/data/dtos"

	"gorm.io/gorm"
)

/*
 * Takes current transaction for calculating count accurately for where queries
 */
func Paginate(pagination *dtos.Pagination, tx *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalElements int64
	tx.Count(&totalElements)
	pagination.TotalElements = totalElements
	totalPages := int(math.Ceil(float64(totalElements) / float64(pagination.Size)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
