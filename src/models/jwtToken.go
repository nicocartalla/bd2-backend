package models

// JwtToken is a JWT token.
type UserProfile struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Major string `json:"major"`
	Role string `json:"role"`
}

type JwtToken struct {
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
	UserProfile UserProfile `json:"user_profile"`
	
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
