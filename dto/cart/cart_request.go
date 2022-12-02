package cartdto

type CreateCart struct {
	UserID int `json:"user_id"`
	// Price     int   `json:"price" form:"price"`
	ProductID int   `json:"product_id" form:"product_id"`
	ToppingID []int `json:"topping_id" form:"topping_id"`
	Qty       int   `json:"qty" form:"qty"`
}

type UpdateCart struct {
	UserID int `json:"user_id"`
	Qty    int `json:"qty" form:"qty"`
}
