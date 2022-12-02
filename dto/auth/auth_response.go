package authdto

type RegisterResponse struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Token    string `json: "token" form: "password" validate: "required"`
}

type LoginResponse struct {
	Fullname string `json: "fullname" form: "fullname" validate: "required"`
	Email    string `json: "email" form: "email" validate: "required"`
	Token    string `json: "token" form : "token" validate: "required"`
}
