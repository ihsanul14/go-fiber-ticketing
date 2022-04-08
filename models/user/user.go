package user

type Request struct {
	UserId   string `json:"user_id" xml:"user_id" form:"user_id"`
	Username string `json:"username" xml:"username" form:"username"`
}

type Response struct {
	UserId     string `json:"user_id" gorm:"user_id"`
	Username   string `json:"username" gorm:"username"`
	Created_at string `json:"created_at" gorm:"created_at"`
	Updated_at string `json:"updated_at" gorm:"updated_at"`
}

type ResponseAll struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Response `json:"data"`
}
