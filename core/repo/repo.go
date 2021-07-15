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
	db, err := sqlx.Connect("mysql", "sql6425032:bnMVGBV6SF@(sql6.freesqldatabase.com:3306)/sql6425032")
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