package database

import "strconv"

type tweetData struct{
	AccountName string
	tweetBofy string
}
func GetTweets() []tweetData{
	var tweetList []tweetData
	for i:=0;i<20;i++{
		tweet := tweetData{
			"@HMSR"+strconv.Itoa(i),
			strconv.Itoa(i)+"カフェオレ",
		}
		tweetList = append(tweetList,tweet)
	}
	return tweetList
}