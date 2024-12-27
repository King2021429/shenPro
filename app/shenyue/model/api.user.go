package model

type UserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
