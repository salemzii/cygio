package facebook

import (
	"errors"
	"fmt"
	"log"
	"os"

	fb "github.com/huandu/facebook/v2"
)

var (
	UnmarshalError = errors.New("Unmarshal Error")
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

func PagePost(text string) {
	res, err := fb.Post("100624668732142/feed",
		fb.Params{
			"message":      text,
			"access_token": os.Getenv("fb_pg_access"),
		},
	)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}

//https://drive.google.com/file/d/1WkVIKJ8Fba6vSDC8DsOnXYQaPBBCvGC6
func PagePostMedia() {
	res, err := fb.Post("100624668732142/photos", fb.Params{
		"url":          "https://drive.google.com/file/d/1WkVIKJ8Fba6vSDC8DsOnXYQaPBBCvGC6",
		"access_token": "EAAK3LCZBTUA4BAAexXPsSd2Cia16VDwVFHWkL0mBblSDIZCFDhwcCVy2fZCBkaLM2p6HxLJhdRjDI3nV6vrXgv97WVdVUSdE3KDdR8HRm9824zLcONGX0NeEhiycPAA28X70S2tbIZAfI4a6qRjrMh785hp3gpgoyWRe6ZAZAVWbKuEvwvvE3ai5DFUcBAYQJhBRh2OKgmt7EtsrJ0HUAZB", //os.Getenv("fb_pg_access"),
	})
	if err != nil {
		log.Println("ERROR: ", err)
	}
	fmt.Println(res)
}

func PageGetPost() {

}

type FbPostResponse struct {
	Id string `json:"id"`
}

func FetchMyFeed() {
	res, err := fb.Get("/me/feed", fb.Params{
		"access_token": os.Getenv("fb_access"),
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
	// read my last feed story.
	fmt.Println("My latest feed story is:", res.Get("data.0.story"))

}

func SearchPage() {
	res, err := fb.Get("/pages/search", fb.Params{
		"access_token": os.Getenv("fb_pg_access"),
		"q":            "golang, postgres",
	})

	if err != nil {
		log.Fatalf("Error fetching data, :%v", err)
	}

	fmt.Println(res)
	var items []fb.Result

	if err = res.DecodeField("data", &items); err != nil {
		fmt.Printf("An error has happened %v", err)
	}

	for _, item := range items {
		fmt.Println(item["id"])
	}
}
