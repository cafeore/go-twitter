package database

import (
	//"strconv"
	"os"
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
   // "net/http"
	//"github.com/labstack/echo"
    //"html/template"
    //"io"
)
func db_exec(db *sql.DB, q string) {
    var _, err = db.Exec(q)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
type tweetData struct{
	AccountName string
	TweetBody string
}
func GetTweets() []tweetData{
	var db *sql.DB
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        fmt.Println(err)
        db.Close()
    }
	var tweetList []tweetData
	hoge,err := db.Query("SELECT * FROM user")
    if err!=nil {
        fmt.Println(err)
    }
    for hoge.Next(){
        var username string
        var password string
        err = hoge.Scan(&username,&password)
        tweet := tweetData{
			username,
			password,
		}
		tweetList = append(tweetList,tweet)
    }
	return tweetList
}