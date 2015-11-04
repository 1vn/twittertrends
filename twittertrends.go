package main

import (
	"fmt"
	"github.com/chimeracoder/anaconda"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type TwitterTrendsCache{
	Woeid int `json: "woeid"`
	TrendData string `json: trendata`
}

func main() {
	anaconda.SetConsumerKey(_CONSUMER_KEY)
	anaconda.SetConsumerSecret(_CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(_APP_TOKEN, _APP_TOKEN_SECRET)
	for {
		searchResult, _ := api.GetTrendsByPlace(1, nil)
		fmt.Println(searchResult.Trends)
		for _, tweet := range searchResult.Trends {
			fmt.Println(tweet.Name)
		}
		time.Sleep(time.Minute * 5)
	}
}
