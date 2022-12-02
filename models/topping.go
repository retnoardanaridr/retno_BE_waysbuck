package models

type Topping struct {
	ID     int          `json:"id" gorm:"primary_key:auto_increment"`
	Title  string       `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int          `json:"price" form:"price" gorm:"type: int"`
	Qty    int          `json:"qty" form:"qty"`
	Image  string       `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"-"`
}

type ToppingCart struct {
	ID    int
	Title string
}
