package _type

import "time"

//Database related type
type User struct {
	ID int `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Username string ` json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Gender string `json:"gender" db:"gender"`
	DateOfBirth string `json:"date_of_birth" db:"dob"`
}

type ResponseUser struct {
	ID int `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Username string ` json:"username" db:"username"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Gender string `json:"gender" db:"gender"`
	DateOfBirth string `json:"date_of_birth" db:"dob"`
}

type Post struct {
	ID int  `json:"id" db:"id"`
	OwnerID int `json:"owner_id" db:"owner_id"`
	Name string `json:"name" db:"name"`
	Brand string `json:"brand" db:"brand"`
	Type string `json:"type" db:"type"`
	Amount int `json:"amount" db:"amount"`
	Description string `json:"description" db:"description"`
	ImageUrl string `json:"image_url" db:"image_url"`
	MainImage string `json:"main_image"`
	AdditionalImage []string `json:"additional_image"`
	Status string `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type NegotiatingPost struct {
	TransactionID int `json:"transaction_id"`
	TradingPost Post `json:"trading_post"`
}

type Transaction struct {
	TransactionID int `json:"transaction_id" db:"id" `
	FromPostID int `json:"from_post_id" db:"from_post_id"`
	ToPostID int `json:"to_post_id" db:"to_post_id"`
	FromUserID int `json:"from_user_id" db:"from_user_id"`
	ToUserID int `json:"to_user_id" db:"to_user_id"`
	FromPostName string `json:"from_post_name"`
	ToPostName string `json:"to_post_name"`
	FromUserName string `json:"from_user_name"`
	ToUserName string `json:"to_user_name"`
	Status string `json:"status" db:"status"`
	Extra string `json:"extra"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
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

type EditPostRequest struct {
	ID int `json:"id"`
	OwnerID int `json:"owner_id"`
	Name string `json:"name"`
	Brand string `json:"brand"`
	Type string `json:"type"`
	Amount int `json:"amount"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
}

type EditPostResponse struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
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

type GetAllUserRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type GetPostByIDRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}
type GetUserByIDRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}


type CreateTransactionRequest struct {
	FromPostID int `json:"from_post_id"`
	ToPostID int `json:"to_post_id"`
}

type CreateTransactionRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type Get_B_BY_A_Res struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type CompleteTransRequest struct {
	TransactionID int `json:"transaction_id"`
}

type CompleteTransRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type SearchRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type DeactivatePostRequest struct {
	PostID int `json:"post_id"`
}

type DeactivatePostRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}

type GetHistoryRes struct {
	StatusCode int `json:"status_code"`
	Detail string `json:"detail"`
}