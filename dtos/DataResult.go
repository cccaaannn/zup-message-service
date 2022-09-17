package dtos

type DataResult[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}
