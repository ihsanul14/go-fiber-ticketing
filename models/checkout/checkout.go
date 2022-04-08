package ticket

type Request struct {
	Id          string `json:"id" xml:"id" form:"id"`
	UserId      string `json:"user_id" xml:"user_id" form:"user_id"`
	TicketId    string `json:"ticket_id" xml:"ticket_id" form:"ticket_id"`
	IsPurchased bool   `json:"is_purchased" xml:"is_purchased" form:"is_purchased"`
}

type Response struct {
	CheckoutId string  `json:"checkout_id" gorm:"checkout_id"`
	UserId     string  `json:"user_id" gorm:"user_id"`
	Username   string  `json:"username" gorm:"username"`
	TicketId   string  `json:"ticket_id" gorm:"ticket_id"`
	Acara      string  `json:"acara" gorm:"acara"`
	Harga      float64 `json:"harga" gorm:"harga"`
	Created_at string  `json:"created_at" gorm:"created_at"`
	Updated_at string  `json:"updated_at" gorm:"updated_at"`
}

type ResponseAll struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	Data       []Response `json:"data"`
	TotalHarga float64    `json:"total_harga"`
}
