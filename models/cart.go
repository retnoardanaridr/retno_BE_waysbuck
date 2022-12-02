package models

import "time"

type Cart struct {
	ID        int       `json:"id" gorm:"primary_key: auto_increment"`
	Qty       int       `json:"qty" gorm:"type: int"`
	Price     int       `json:"price" gorm:"type: int"`
	ProductID int       `json:"product_id"`
	Product   Product   `json:"product" gorm:"-"`
	ToppingID []int     `json:"topping_id" form:"topping_id" gorm:"-"`
	Topping   []Topping `json:"topping" gorm:"-"`
	BuyerID   int       `json:"buyer_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartResponse struct {
	ID        int     `json: "id"`
	UserID    int     `json:"user_id"`
	Qty       int     `json: "qty"`
	price     int     `json: "price"`
	ProductID int     `json: "product_id`
	Product   Product `json:"product"`
	ToppingID []int   `json:"topping_id"`
	Topping   []int   `json: "-"`
}

func (CartResponse) TableName() string {
	return "carts"
}
