package cartdto

type CartResponse struct {
	ID       int `json: "id"`
	UserID   int `json:"user_id"`
	Subtotal int `json:"sub_total"`
}
