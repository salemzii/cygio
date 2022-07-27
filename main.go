package main

import (
	"fmt"

	"github.com/salemzii/cygio/app"
)

func main() {
	fmt.Println(app.GetTweetById(1552302570470883329))
	fmt.Println(app.CreateTweet("Hello twitter, hope this works :) "))
}
