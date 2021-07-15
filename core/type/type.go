package _type

//Database related type
type User struct {
	ID int `db:"id"`
	Email string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Gender string `db:"gender"`
	DateOfBirth string `db:"dob"`
}

//Request, response type
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type SignupRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Gender string `json:"gender"`
	DOB string `json:"dob"`
}

type SignupResponse struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}


