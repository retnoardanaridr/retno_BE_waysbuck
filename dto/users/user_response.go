package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form: "email" validate: "required"`
	Password string `json:"-"`
	Role     string `json: "role"`
	Token    string `json: "token" form: "password" validate: "required"`
}

func (UserResponse) TableName() string {
	return "users"
}
