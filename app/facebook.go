package app

import (
	"fmt"
	"log"
	"os"

	fb "github.com/huandu/facebook/v2"
)

func GetUsername() {
	res, err := fb.Get("/me", fb.Params{
		"fields":       "first_name",
		"access_token": os.Getenv("fb_access"),
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Hello", res["first_name"])
}

func PagePost() {
	res, err := fb.Post("100624668732142/feed",
		fb.Params{
			"message":      "hello world from a gopher 14:53",
			"access_token": os.Getenv("fb_pg_access"),
		},
	)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}

type FbPostResponse struct {
	Id string `json:"id"`
}
