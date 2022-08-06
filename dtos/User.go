package dtos

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
