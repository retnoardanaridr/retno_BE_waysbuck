package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Image     string    `json: "image" gorm: "type: varchar(255)"`
	Role      string    `json: "role"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserResponse struct {
	ID       int    `json: "id"`
	Fullname string `json: "fullname"`
	Email    string `json: "email"`
	Image    string `json: "image"`
	Role     string `json: "role"`
}

func (UserResponse) TableName() string {
	return "users"
}
