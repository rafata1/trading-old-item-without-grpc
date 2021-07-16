package _type

import "time"

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

type Post struct {
	ID int  `json:"id "db:"id"`
	OwnerID int `json:"owner_id" db:"owner_id"`
	Name string `json:"name" db:"name"`
	Brand string `json:"brand" db:"brand"`
	Type string `json:"type" db:"type"`
	Amount int `json:"amount" db:"amount"`
	Description string `json:"description" db:"description"`
	ImageURL string `json:"image_url" db:"image_url"`
	Status string `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"create_at"`
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

type AddPostRequest struct {
	OwnerID int `json:"owner_id"`
	Name string `json:"name"`
	Brand string `json:"brand"`
	Type string `json:"type"`
	Amount int `json:"amount"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
}

type AddPostResponse struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type GetUserPostRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type GetAllPostRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type GetPostByIDRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}
