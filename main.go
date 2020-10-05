package main

import(
	// "config"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
)

type CommentModel struct{
	Db *sql.DB
}

type LikeModel struct{
	Db *sql.DB
}

type Comment struct{
	id_comment int64
	comment string
}

type Like struct{
	id_like int64
	like string
}

type response struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Data int
}

func GetDB()(db *sql.DB, err error){
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/slivth-login")

	return
}

func (commentModel CommentModel) Count() (int64, error){
	rows, err := commentModel.Db.Query("SELECT COUNT(*) AS count_comment FROM comment")
	if err != nil{
		return 0, err
	} else{
		var count_comment int64
		for rows.Next(){
			rows.Scan(&count_comment)
		}
		return count_comment, nil
	}
}

func EndCount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET"{
		var response response

		db, err := GetDB()
		if err != nil{
			fmt.Println(err)
		} else {
			commentModel := CommentModel{
			Db: db,
			}
			countComment, err2 := commentModel.Count()
			if err2 != nil{
				fmt.Println(err2)
			} else {
				fmt.Println("Count Comment:", countComment)
				// var data = (countComment)
				response.Status = 1
				response.Message = "Success Count Comment"
				// response.Data = data
			}
			json.NewEncoder(w).Encode(countComment)
		}
	}
}

func (likeModel LikeModel) Count() (int64, error){
	rows, err := likeModel.Db.Query("SELECT COUNT(*) AS count_like FROM likes")
	if err != nil{
		return 0, err
	} else{
		var count_like int64
		for rows.Next(){
			rows.Scan(&count_like)
		}
		return count_like, nil
	}
}

func EndLike(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET"{
		var response response

		db, err := GetDB()
		if err != nil{
			fmt.Println(err)
		} else {
			likeModel := LikeModel{
			Db: db,
			}
		countLike, err2 := likeModel.Count()
		if err2 != nil{
			fmt.Println(err2)
		} else {
			fmt.Println("Count Like:", countLike)
		 	// var data = (countComment)
			response.Status = 1
			response.Message = "Success Count Likes"
			// response.Data = data
			}
			json.NewEncoder(w).Encode(countLike)
		}
	}
}

func main(){
	http.HandleFunc("/CountComment", EndCount)
	http.HandleFunc("/CountLike", EndLike)
	fmt.Println("Server running on port:8080")
	http.ListenAndServe(":8080", nil)
}