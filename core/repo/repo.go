package repo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	_type "trading.olditem.app/core/type"
)

var myDB *sqlx.DB

func ConnectToDB() {
	db, err := sqlx.Connect("mysql", "sql6425032:bnMVGBV6SF@(sql6.freesqldatabase.com:3306)/sql6425032?parseTime=true")
	myDB = db
	if err != nil {
		log.Fatalf("Err occured while connecting to db")
	}
	fmt.Println("Connect to db successfully")
}

func SignUp(email string, username string, password string, phone_number string,
	gender string, dob string) (statusCode int,
	detail string) {

	//check if email is already existed
	var user _type.User
	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE email=? LIMIT 1", email)

	if user.Email == email {
		return 101, "Email đã tồn tại"
	}

	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE username=? LIMIT 1", username)

	if user.Username == username {
		return 202, "Username đã tồn tại"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	queryCode := `INSERT INTO Users (email, username, password, phone_number, gender, dob) VALUES (?, ?, ?, ?, ?, ?)`
	myDB.MustExec(queryCode, email, username, string(hashedPassword), phone_number, gender, dob)
	return 200, "Thêm tài khoản thành công"
}

func Login(email string, password string) (statusCode int, detail string) {

	//check if email is existed
	var user _type.User
	myDB.Get(&user, "SELECT id, email, username, password, phone_number, gender, dob FROM Users WHERE email=? LIMIT 1", email)
	if user.Email != email {
		return 101, "Email không tồn tại"
	}
	//check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 202, "Password sai"
	}

	return 200, "Đăng nhập thành công"
}

func AddPost(owner_id int, name string, brand string, _type string, amount int, description string,
	image_url string) (statusCode int, detail string) {

	queryCode := `INSERT INTO posts (owner_id, name, brand, type, amount, description, image_url, status) VALUES (?, 
?, ?, ?, ? , ? ,? ,? )`
	myDB.MustExec(queryCode, owner_id, name, brand, _type, amount, description, image_url, "available")
	return 200, "Thêm bài thành công"
}

func GetAllPost() (statusCode int, detail string, result *[]_type.Post) {
	posts := []_type.Post {}
	err := myDB.Select(&posts,`SELECT * FROM posts WHERE status = "available"`)
	if err != nil {
		return 101, "Cannot get posts from db", &[]_type.Post{}
	}
	return  200, "Lấy dữ liệu thành công", &posts
}

func GetPostByID(post_id int) (statusCode int, detail string, result *_type.Post) {

	post := _type.Post{}
	err := myDB.Get(&post, "SELECT * FROM posts WHERE id = ?", post_id)
	if err != nil {
		return 101, "Cannot get  post by id", &_type.Post {}
	}

	return 200, "Lấy dữ liệu thành công", &post

}

func GetPostsOfUser(owner_id int) (statusCode int, detail string, result *[]_type.Post) {
	userPosts := []_type.Post {}
	err := myDB.Select(&userPosts,`SELECT * FROM posts WHERE owner_id = ?`,owner_id)
	if err != nil {
		panic(err)
		return 101, "Cannot get user's posts", &[]_type.Post {}
	}
	return 200, "Lấy bài thành công", &userPosts
}