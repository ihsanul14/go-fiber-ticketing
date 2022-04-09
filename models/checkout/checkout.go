package ticket

type Request struct {
	Id             int     `json:"id" xml:"id" form:"id"`
	CheckoutId     string  `json:"checkout_id" xml:"checkout_id" form:"checkout_id"`
	UserId         string  `json:"user_id" xml:"user_id" form:"user_id"`
	TicketId       string  `json:"ticket_id" xml:"ticket_id" form:"ticket_id"`
	PaymentAccount float64 `json:"payment_account" xml:"payment_account" form:"payment_account"`
	IsPurchased    bool    `json:"is_purchased" xml:"is_purchased" form:"is_purchased"`
}

type CreateRequest struct {
	CheckoutId string   `json:"checkout_id" xml:"checkout_id" form:"checkout_id"`
	UserId     string   `json:"user_id" xml:"user_id" form:"user_id"`
	TicketId   []string `json:"ticket_id" xml:"ticket_id" form:"ticket_id"`
}

type Response struct {
	Id          int    `json:"id" gorm:"id"`
	CheckoutId  string `json:"checkout_id" gorm:"checkout_id"`
	UserId      string `json:"user_id" gorm:"user_id"`
	TicketId    string `json:"ticket_id" gorm:"ticket_id"`
	IsPurchased bool   `json:"is_purchased" gorm:"is_purchased"`
	Created_at  string `json:"created_at" gorm:"created_at"`
	Updated_at  string `json:"updated_at" gorm:"updated_at"`
}

type ResponseSummary struct {
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
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Response `json:"data"`
}

type ResponseSummaryAll struct {
	Code       int               `json:"code"`
	Message    string            `json:"message"`
	Data       []ResponseSummary `json:"data"`
	TotalHarga float64           `json:"total_harga"`
}
