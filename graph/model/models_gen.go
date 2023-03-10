// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type LinkInput struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
