package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	AvatarLink string `json:"avatar"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
