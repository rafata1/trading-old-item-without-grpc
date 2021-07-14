package _type

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}
