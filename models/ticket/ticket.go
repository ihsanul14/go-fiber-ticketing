package ticket

type Request struct {
	TicketId string  `json:"ticket_id" xml:"ticket_id" form:"ticket_id"`
	Acara    string  `json:"acara" xml:"acara" form:"acara"`
	Harga    float64 `json:"harga" xml:"harga" form:"harga"`
}

type Response struct {
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
