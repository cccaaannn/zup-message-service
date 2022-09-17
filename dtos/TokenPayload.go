package dtos

type TokenPayload struct {
	Id         uint64 `json:"id"`
	UserStatus uint64 `json:"userStatus"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	TokenType  uint64 `json:"tokenType"`
}
