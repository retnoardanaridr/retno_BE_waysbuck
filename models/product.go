package models

type Product struct {
	ID     int          `json:"id" gorm:"primary_key:auto_increment"`
	Title  string       `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int          `json:"price" form:"price" gorm:"type: int"`
	Image  string       `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int          `json:"user_id"`
	Qty    int          `json:"qty" form:"qty" gorm:"type: varchar(255)"`
	User   UserResponse `json:"-"`
}

type ProductResponse struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Title  string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"-"`
	Qty    int
}

func (ProductResponse) TableName() string {
	return "products"
}
