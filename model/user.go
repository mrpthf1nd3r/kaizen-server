package model

type User struct {
	Email    string
	Password string
}

type OutUser struct {
	Email string `json:"email"`
}

func (user *User) Out() *OutUser {
	return &OutUser{
		Email: user.Email,
	}
}
