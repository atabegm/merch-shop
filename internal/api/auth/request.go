package auth

// Request object create.
type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
