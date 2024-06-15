package models
import (
	"time"
)
// JwtToken is a JWT token.
type UserProfile struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Major string `json:"major"`
	RoleID int `json:"role_id"`
}

type JwtToken struct {
	Token      string `json:"token"`
	Expiration time.Time  `json:"expiration"`
	UserProfile UserProfile `json:"user_profile"`
	
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
