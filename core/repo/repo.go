package repo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	_type "trading.olditem.app/core/type"
	"unicode"
)

var myDB *sqlx.DB

func ExtractImage(image_url string) (string, []string) {
	temp := strings.Split(image_url, " ")
	return temp[0], temp[1:]
}

func ConnectToDB() bool {
	db, err := sqlx.Connect("mysql", "admin:2XXCdB2z@(mysql-40280-0.cloudclusters.net:280)/trading_web_data?parseTime=true")
	myDB = db
	if err != nil {
		fmt.Println("Err occured while connecting to db")
		return false
	}
	fmt.Println("Connect to db successfully")
	return true
}

func SignUp(email string, username string, password string, phone_number string,
	gender string, dob string) (statusCode int,
	detail string) {

	//check if email is already existed
	var user _type.User
	myDB.Get(&user, "SELECT id, email, username, password FROM users WHERE email=? LIMIT 1", email)

	if user.Email == email {
		return 101, "Email đã tồn tại"
	}

	myDB.Get(&user, "SELECT id, email, username, password FROM users WHERE username=? LIMIT 1", username)

	if user.Username == username {
		return 202, "Username đã tồn tại"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	queryCode := `INSERT INTO users (email, username, password, phone_number, gender, dob) VALUES (?, ?, ?, ?, ?, ?)`
	myDB.MustExec(queryCode, email, username, string(hashedPassword), phone_number, gender, dob)
	return 200, "Thêm tài khoản thành công"
}

func Login(email string, password string) (statusCode int, detail string, result *_type.ResponseUser) {

	//check if email is existed
	var user _type.User
	myDB.Get(&user, "SELECT id, email, username, password, phone_number, gender, dob FROM users WHERE email=? LIMIT 1", email)
	if user.Email != email {
		return 101, "Email không tồn tại", &_type.ResponseUser{}
	}
	//check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 202, "Password sai", &_type.ResponseUser{}
	}

	return 200, "Đăng nhập thành công", &_type.ResponseUser{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		PhoneNumber: user.PhoneNumber,
		Gender: user.Gender,
		DateOfBirth: user.DateOfBirth,
	}
}

func AddPost(owner_id int, name string, brand string, _type string, amount int, description string,
	image_url string) (statusCode int, detail string) {

	queryCode := `INSERT INTO posts (owner_id, name, brand, type, amount, description, image_url, status) VALUES (?, 
?, ?, ?, ? , ? ,? ,? )`
	myDB.MustExec(queryCode, owner_id, name, brand, _type, amount, description, image_url, "available")
	return 200, "Thêm bài thành công"
}

func EditPost(id int,owner_id int, name string, brand string, _type string, amount int, description string,
	image_url string) (statusCode int, detail string) {
	queryCode := `UPDATE posts SET owner_id=?, name=?, brand=?, type=?, amount=?, description=?, image_url=?, status="available" WHERE id = ? `
	_, err := myDB.Exec(queryCode, owner_id, name, brand, _type, amount, description, image_url, id)
	if err != nil {
		return 101, "Cập nhật sản phẩm thất bại"
	}
	return 200, "Cập nhật sản phẩm thành công"
}

func DeactivatePost(post_id int) (statusCode int, detail string)  {

	queryCode := `UPDATE posts SET status = "unavailable" WHERE id = ?`
	myDB.MustExec(queryCode, post_id)

	// deactivate all transactions related to the post
	queryCode = `UPDATE transactions SET status = "rejected" WHERE from_post_id =  ? OR to_post_id = ?`
	myDB.MustExec(queryCode, post_id, post_id)

	return 200, "Đã vô hiệu hóa bài viết"

}

func GetAllPost() (statusCode int, detail string, result *[]_type.Post) {
	posts := []_type.Post {}
	err := myDB.Select(&posts,`SELECT * FROM posts WHERE status = "available"`)
	if err != nil {
		return 101, "Cannot get posts from db", &[]_type.Post{}
	}

	for i:=0; i < len(posts); i++ {
		posts[i].MainImage, posts[i].AdditionalImage = ExtractImage(posts[i].ImageUrl)
	}

	return  200, "Lấy dữ liệu thành công", &posts
}

func GetAllUser() (statusCode int, detail string, result *[]_type.ResponseUser) {
	users := []_type.ResponseUser{}
	err := myDB.Select(&users, `SELECT id, email, username, phone_number, gender, dob FROM users` )
	if err != nil {
		return 101, "Cannot get users from db", &[]_type.ResponseUser{}
	}
	return 200, "Lấy dữ liệu thành công", &users
}

func GetUserByID(user_id int) (statusCode int, detail string, result *_type.ResponseUser) {
	user := _type.ResponseUser{}
	err := myDB.Get(&user, `SELECT id, email, username, phone_number, gender, dob FROM users WHERE id = ? LIMIT 1`, user_id )
	if err != nil {
		return 101, "Cannot get user by id, may be user_id is not existed", &_type.ResponseUser{}
	}
	return 200, "Lấy dữ liệu thành công", &user
}

func GetPostByID(post_id int) (statusCode int, detail string, result *_type.Post) {

	post := _type.Post{}
	err := myDB.Get(&post, "SELECT * FROM posts WHERE id = ?", post_id)
	if err != nil {
		return 101, "Cannot get  post by id", &_type.Post {}
	}
	post.MainImage, post.AdditionalImage = ExtractImage(post.ImageUrl)
	return 200, "Lấy dữ liệu thành công", &post

}

func GetPostsOfUser(owner_id int) (statusCode int, detail string, result *[]_type.Post) {
	userPosts := []_type.Post {}
	err := myDB.Select(&userPosts,`SELECT * FROM posts WHERE owner_id = ? AND status = "available"`,owner_id)
	if err != nil {
		return 101, "Cannot get user's posts", &[]_type.Post {}
	}

	for i:=0; i < len(userPosts); i++ {
		userPosts[i].MainImage, userPosts[i].AdditionalImage = ExtractImage(userPosts[i].ImageUrl)
	}

	return 200, "Lấy bài thành công", &userPosts
}

func CreateTransaction(fromPostId int, toPostId int) (statusCode int, detail string) {

	//check if transaction existed
	trans_id := -1
	_ = myDB.Get(&trans_id, `SELECT id FROM transactions WHERE from_post_id = ? AND to_post_id = ? LIMIT 1`, fromPostId, toPostId)

	if trans_id != -1 {
		return 101, "Giao dịch này đã tồn tại rồi"
	}

	// update user_history
	var from_user_id int
	_ = myDB.Get(&from_user_id, `SELECT owner_id FROM posts WHERE id = ?`, fromPostId)
	var to_user_id int
	_ = myDB.Get(&to_user_id, `SELECT  owner_id FROM posts WHERE id = ?`, toPostId)

	status := "negotiating"
	queryCode := `INSERT INTO transactions (from_post_id, from_user_id, to_post_id, to_user_id, status) VALUES (?, ?, ?, ?, ?)`
	myDB.MustExec(queryCode, fromPostId, from_user_id, toPostId, to_user_id, status)

	return 200, "Thêm transaction thành công"

}

func  Get_B_By_A(toPostId int) (statusCode int, detail string, result *[]_type.NegotiatingPost) {

	res := []_type.NegotiatingPost{}
	dbPost := []_type.Post{}

	err := myDB.Select(&dbPost,`SELECT * FROM posts WHERE id IN (SELECT from_post_id FROM transactions WHERE to_post_id = ? AND status = "negotiating")`, toPostId)

	if err != nil {
		return 101, "query DB failed", &[]_type.NegotiatingPost{}
	}

	for _, p := range dbPost {
		var transID int
		err := myDB.Get(&transID, `SELECT id FROM transactions WHERE from_post_id = ? AND to_post_id = ?`, p.ID, toPostId)
		if err != nil {
			panic(err)
			return 201, "query transid related to from_post_id failed", &[]_type.NegotiatingPost{}
		}
		res = append(res, _type.NegotiatingPost{
			TransactionID: transID,
			TradingPost: p,
		})
	}

	for i:=0; i< len(res); i++ {
		res[i].TradingPost.MainImage, res[i].TradingPost.AdditionalImage = ExtractImage(res[i].TradingPost.ImageUrl)
	}

	return 200, "Lấy bài thành công", &res
}

func CompleteTransaction(trans_ID int) (statusCode int, detail string)  {

	var from_post_id int
	err := myDB.Get(&from_post_id, "SELECT from_post_id from transactions WHERE id = ?", trans_ID)
	if err != nil {
		return 101, "update db failed"
	}

	var to_post_id int
	err = myDB.Get(&to_post_id, "SELECT to_post_id from transactions WHERE id = ?", trans_ID)
	if err != nil {
		return 101, "update db failed"
	}

	DeactivatePost(from_post_id)
	DeactivatePost(to_post_id)

	_, err = myDB.Exec(`UPDATE  transactions SET status = "Completed" WHERE id = ?`, trans_ID )

	if err != nil {
		return 101, "update db failed"
	}

	return 200, "updated"
}

func normalizeUnicodeString(keywords string) string  {
	trans := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(trans, keywords)
	result = strings.ReplaceAll(result, "đ", "d")
	result = strings.ReplaceAll(result, "Đ", "D")
	return strings.ToLower(result)
}

func getKeyWordOfPost(post _type.Post) string {
	// Namex4, typex3, brandx2
	return normalizeUnicodeString(post.Name) + " " + normalizeUnicodeString(post.Name) + " " + normalizeUnicodeString(post.Name) + " " + normalizeUnicodeString(post.Name)  + " "+ normalizeUnicodeString(post.Brand) + " " + normalizeUnicodeString(post.Brand) + " " + normalizeUnicodeString(post.Type) + " " + normalizeUnicodeString(post.Type) + " " + normalizeUnicodeString(post.Type)
}

func computeMatchPoint(keyword string, post_keyword string) int {

	point :=0
	a := strings.Split(keyword, " ")
	b := strings.Split(post_keyword, " ")

	for _,x := range a {
		for _,y := range b {
			if x == y {
				point ++
			}
		}
	}

	return point
}


func SearchHardCore(keywords string) (statusCode int, detail string, result *[]_type.Post) {

	keywords = normalizeUnicodeString(keywords)
	dbPosts := []_type.Post {}
	err := myDB.Select(&dbPosts,`SELECT * FROM posts WHERE status = "available"`)
	res := []_type.Post{}

	if err != nil {
		return 101, "Cannot get posts from db", &[]_type.Post{}
	}

	type Couple struct {
		ORD int
		Point int
	}

	tmp := []Couple{}

	for i,p := range dbPosts {
		tmp = append(tmp, Couple{ORD: i, Point: computeMatchPoint(keywords, getKeyWordOfPost(p)) })
	}

	//sort tmp by match point
	for i:= 0; i < len(tmp); i++ {
		for j:= i + 1; j < len(tmp); j++ {
			if tmp[i].Point < tmp[j].Point {
				c := tmp[i]
				tmp[i] = tmp[j]
				tmp[j] = c
			}
		}
	}

	for _,c := range tmp {
		res = append(res, dbPosts[c.ORD])
	}

	for i:=0; i<len(res); i++ {
		res[i].MainImage, res[i].AdditionalImage = ExtractImage(res[i].ImageUrl)
	}

	return  200, "Lấy dữ liệu thành công", &res
}

func GetUserHistory(userID int) (statusCode int, detail string, result *[]_type.Transaction) {

	res := []_type.Transaction{}
	err := myDB.Select(&res, `SELECT * FROM transactions WHERE (from_user_id = ? OR to_user_id = ?) AND (status = "Completed" OR  status = "rejected") `, userID, userID)
	if err != nil {
		panic(err)
		return 101, "Lỗi truy vấn db", &[]_type.Transaction{}
	}

	for i:=0; i < len(res); i++ {
		var fromUserName string
		var toUserName string
		var fromPostName string
		var toPostName string

		_ = myDB.Get(&fromUserName, "SELECT username FROM users WHERE id = ?", res[i].FromUserID)
		_ = myDB.Get(&toUserName, "SELECT username FROM users WHERE id = ?", res[i].ToUserID)
		_ = myDB.Get(&fromPostName, "SELECT name FROM posts WHERE id = ?", res[i].FromPostID)
		_ = myDB.Get(&toPostName, "SELECT  name FROM posts WHERE id = ?", res[i].ToPostID)
		res[i].FromUserName = fromUserName
		res[i].ToUserName = toUserName
		res[i].FromPostName = fromPostName
		res[i].ToPostName = toPostName

		if res[i].Status == "Completed" {
			res[i].Extra = "Bạn đã trao đổi thành công giao dịch này"
		} else
		{
			if userID == res[i].FromUserID {
				res[i].Extra = "Bạn đã bị từ chối giao dịch này"
			} else {
				res[i].Extra = "Bạn đã từ chối giao dịch này"
			}
		}
	}
	return 200, "Thành công", &res
}