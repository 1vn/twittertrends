package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/chimeracoder/anaconda"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type TwitterTrendsCache struct {
	Woeid     int    `json: "woeid"`
	TrendData string `json: "trendata"`
}

func (cache TwitterTrendsCache) String() string {
	return fmt.Sprintf("%s|%s", cache.Woeid, cache.TrendData)
}

func main() {
	anaconda.SetConsumerKey(_CONSUMER_KEY)
	anaconda.SetConsumerSecret(_CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(_APP_TOKEN, _APP_TOKEN_SECRET)
	db, err := sql.Open("mysql", _DB_USER+":"+_DB_PASS+"@/"+_DB_NAME)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		searchResult, _ := api.GetTrendsByPlace(1, nil)
		fmt.Println(searchResult.Trends)
		search, err := json.Marshal(searchResult)
		if err != nil {
			panic(err)
		}
		trends := TwitterTrendsCache{TrendData: string(search), Woeid: 1}
		stmtIns, err := db.Prepare("INSERT INTO twittertrends VALUES(" + trends.TrendData + ")")
		defer stmtIns.Close()
		time.Sleep(time.Minute * 5)
	}
}
