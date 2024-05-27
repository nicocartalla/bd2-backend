package models

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Major      string `json:"major"`
	Password   string `json:"password"`
	Role 	   string `json:"role"`
}

