package main

import (
	"fmt"
	"github.com/chimeracoder/anaconda"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	anaconda.SetConsumerKey(_CONSUMER_KEY)
	anaconda.SetConsumerSecret(_CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(_APP_TOKEN, _APP_TOKEN_SECRET)
	searchResult, _ := api.GetTrendsByPlace(1, nil)
	for _, tweet := range searchResult.Trends {
		fmt.Println(tweet.Name)
	}
}
