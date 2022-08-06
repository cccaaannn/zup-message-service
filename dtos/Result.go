package dtos

type Result struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    *User  `json:"data"`
}
