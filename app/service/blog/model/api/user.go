package api

type User struct {
	Nick     string `json:"nick" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
	QQ       int64  `json:"qq"`
	Wechat   string `json:"wechat"`
	Phone    int64  `json:"phone"`
}
