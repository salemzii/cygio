package twitter

import (
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
)

var cred *Credentials

func init() {
	cred = &Credentials{
		ConsumerKey:       os.Getenv("TWITTER_SECRET_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_SECRET"),
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
	}

}

// Credentials stores all access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// return http.Client types to automatically authorize http.Request's
func (creds *Credentials) GetClientToken() (client *http.Client) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	return config.Client(oauth1.NoContext, token)
}

type Createtweetrequest struct {
	Text string `json:"text,omitempty"`
}

// CreateTweetData is the data returned when creating a tweet
type CreateTweetData struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type DeleteTweetData struct {
	Deleted bool `json:"deleted"`
}
type CreateTweetResponse struct {
	Tweet *CreateTweetData `json:"data"`
}

type DeleteTweetResponse struct {
	Data *DeleteTweetData `json:"data"`
}

type GetTweetResponse struct {
	Data []CreateTweetData `json:"data"`
}
type TwitterUsers struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
type GetLikingUsers struct {
	Data []TwitterUsers `json:"data"`
}
