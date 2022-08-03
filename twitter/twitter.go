package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var (
	httpClient *http.Client
	URL        = "https://api.twitter.com/2"
	Wg         sync.WaitGroup
)

func GetTweetById(id int) *GetTweetResponse {
	endpoint := fmt.Sprintf("/tweets?ids=%d", id)

	httpClient = cred.GetClientToken()

	path := URL + endpoint
	resp, err := httpClient.Get(path)
	if err != nil {
		log.Println("Error: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: ", err)
	}
	var results GetTweetResponse
	if err := json.Unmarshal(body, &results); err != nil {
		log.Println(err)
	}

	return &results
}

func CreateTweet(text string) (response *CreateTweetResponse, err error) {
	endpoint := "/tweets"
	path := URL + endpoint
	tweet := Createtweetrequest{
		Text: text,
	}
	data, err := json.Marshal(&tweet)
	if err != nil {
		log.Println("Error encoding to json", err)
	}

	httpClient = cred.GetClientToken()
	res, err := httpClient.Post(path, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error making request: ", err)
	}

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resp CreateTweetResponse

	if err := json.Unmarshal(respByte, &resp); err != nil {
		log.Println("Error decoding response::: ", err)
	}

	return &resp, nil
}

func DeleteTweet(TweetId int) (response *DeleteTweetResponse, err error) {
	endpoint := fmt.Sprintf("/tweets/%d", TweetId)
	path := URL + endpoint

	httpClient = cred.GetClientToken()
	request, err := http.NewRequest("delete", path, nil)
	if err != nil {
		log.Println("Error making request")
	}
	res, err := httpClient.Do(request)
	if err != nil {
		log.Println("Error making request: ", err)
	}

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resp DeleteTweetResponse

	if err := json.Unmarshal(respByte, &resp); err != nil {
		log.Println("Error decoding response::: ", err)
	}

	return &resp, nil
}

func TweetLikes(TweetId int) (data *GetLikingUsers, err error) {
	endpoint := fmt.Sprintf("/tweets/%d/liking_users", TweetId)
	path := URL + endpoint

	httpClient = cred.GetClientToken()

	res, err := httpClient.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(res.Status)

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respByte))
	var resp GetLikingUsers

	if err := json.Unmarshal(respByte, &resp); err != nil {
		log.Println("Error decoding response::: ", err)
	}

	fmt.Println(resp)
	return &resp, nil
}

func LikedTweets(UserId int) {
	endpoint := fmt.Sprintf("/users/%d/liked_tweets", UserId)
	path := URL + endpoint

	httpClient = cred.GetClientToken()

	res, err := httpClient.Get(path)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respByte))
}

func GetUserByUserName(username string) (response *GetLikingUsers, err error) {
	endpoint := fmt.Sprintf("/users/by?usernames=%s", username)
	path := URL + endpoint

	httpClient = cred.GetClientToken()

	res, err := httpClient.Get(path)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respByte))
	var resp GetLikingUsers

	if err := json.Unmarshal(respByte, &resp); err != nil {
		log.Println("Error decoding response::: ", err)
	}

	fmt.Println(resp)
	return &resp, nil
}

func RecentTweetSearch(keyword string) {

	//https://api.twitter.com/2/tweets/search/recent?
	//query=from:TwitterDev&tweet.fields=created_at&expansions=author_id&user.fields=created_at
	endpoint := fmt.Sprintf("/tweets/search/recent?query=%s&tweet.fields=created_at&expansions=author_id&user.fields=created_at", keyword)
	path := URL + endpoint

	httpClient = cred.GetClientToken()
	res, err := httpClient.Get(path)

	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respByte))
}

func GetVolumeTweet() {
	endpoint := "/tweets/sample/stream"
	path := URL + endpoint

	httpClient = cred.GetClientToken()
	res, err := httpClient.Get(path)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respByte))
}
