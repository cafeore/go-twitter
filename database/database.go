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
    hoge.Close()
	return tweetList
}
func GetUser(name string,pass string) bool{
    var db *sql.DB
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        fmt.Println(err)
        db.Close()
    }
    hoge,err := db.Query("SELECT * FROM user")
    if err!=nil {
        fmt.Println(err)
    }
    for hoge.Next(){
        var username string
        var password string
        err = hoge.Scan(&username,&password)
        if username == name && password == pass{
            hoge.Close()
            return true
        }
    }
    hoge.Close()
    return false
}
func CheckUser(name string,pass string) bool{
    var db *sql.DB
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        fmt.Println(err)
        db.Close()
    }
    hoge,err := db.Query("SELECT * FROM user")
    if err!=nil {
        fmt.Println(err)
    }
    for hoge.Next(){
        var username string
        var password string
        err = hoge.Scan(&username,&password)
        if username == name{
            hoge.Close()
            return true
        }
    }
    hoge.Close()
    return false
}