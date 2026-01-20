package dto

type RegisterRequest struct {
	Username 		string `json:"username"`
	Password 		string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Role     		string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token,omitempty"`

}
