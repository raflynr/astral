package auth

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
}

type Register struct {
	ID                   string
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required,min=8,eqfield=Password"`
	Username             string `json:"username" validate:"required"`
}

type User struct {
	ID       string
	Email    string
	Password string
	Username string
}
