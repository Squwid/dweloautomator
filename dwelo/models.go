package dwelo

import "github.com/google/uuid"

// LoginRequest is the request that gets sent to retrieve a token from dwelo
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AppID    string `json:"applicationId"`

	id string
}

// LoginResponse is the response that comes back from the dwelo login
type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`

	id       string
	loggedIn bool
}

// User is the structure that comes back from dwelo when it returns requests
type User struct {
	UID       int    `json:"uid"`
	FirstName string `json:"firstName"`
	// Other things that dwelo returns:
	// dateRegistered, phone, lastname, but i dont need any of that
}

// NewLoginRequest is the constructor for LoginRequest
func NewLoginRequest(email, password string) *LoginRequest {
	return &LoginRequest{
		Email:    email,
		Password: password,
		AppID:    "concierge",
		id:       uuid.New().String(),
	}
}

// Command is the structure of what happens in dwelo
type Command struct {
	Command string `json:"command"`
	Value   int    `json:"comandValue,omitempty"`
}
