package database

import (
	//"strconv"
	"os"
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
  // "github.com/PuerkitoBio/goquery"
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
    TweetTime string
    Good int
  TweetId int
}
var user string
func GetTweets() []tweetData{
	var db *sql.DB
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        fmt.Println(err)
        db.Close()
    }
	var tweetList []tweetData
	hoge,err := db.Query("SELECT * FROM tweets order by time desc limit 20")
    if err!=nil {
        fmt.Println(err)
    }
    for hoge.Next(){
        var username string
        var body string
        var time string
        var good int
        var tweetid int
        err = hoge.Scan(&username,&body,&time,&good,&tweetid)
        tweet := tweetData{
			      username,
			      body,
            time,
            good,
            tweetid,
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
        var id int
        err = hoge.Scan(&id,&username,&password)
        if username == name && password == pass{
            user = username
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
    if hoge == nil {
        return false
    }
    for hoge.Next(){
        var username string
        var password string
        var id int
        err = hoge.Scan(&id,&username,&password)
        if username == name{
            fmt.Println(user);
            hoge.Close()
            return true
        }
    }
    hoge.Close()
    MakeAccount(name,pass)
    return false
}
func MakeTweet(body string) {
    if body != ""{
        var db *sql.DB
        db, err := sql.Open("sqlite3", "./data.db")
        if err != nil {
            fmt.Println(err)
            db.Close()
        }
        fmt.Println(user);
        var q = ""
        q = "INSERT INTO tweets "
        q+= "( userid,body)"
        q+= " VALUES"
        q+= " (\""+user+"\",\""+body+"\")"
        db_exec(db,q)
    }
   /* doc, err := goquery.NewDocument("https://github.com/PuerkitoBio/goquery")
      if err != nil {
          fmt.Print("url scarapping failed")
      }
      res, err := doc.Find("body").Html()
      if err != nil {
          fmt.Print("dom get failed")
      }
      for i:=0;i<len(res);i++{
        fmt.Println(res[i])
      }*/
}
func MakeAccount(name string , pass string){
    var db *sql.DB
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        fmt.Println(err)
        db.Close()
    }
    var q = ""
    q = "INSERT INTO user "
    q+= " (name,pass)"
    q+= " VALUES"
    q+= " (\""+name+"\",\""+pass+"\")"
    db_exec(db,q)
}
func  GoodPlus(tweetId string) {
    var db *sql.DB
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        fmt.Println(err)
        db.Close()
    }
    var q = ""
    q = "UPDATE tweets SET good = good + 1 WHERE id ="
    q+= tweetId + ";"
    fmt.Println(q)
    db_exec(db,q)
}
