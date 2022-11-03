package dtos

type ListDataResult[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    *[]T   `json:"data"`
}
