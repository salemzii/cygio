package linkedin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LinkedInError struct {
	Message          string `json:"message"`
	ServiceErrorCode int    `json:"serviceErrorCode"`
	Status           int    `json:"status"`
}
type AuthResponse struct {
	Access_Token string `json:"access_token"`
	Expires_In   string `json:"expires_in"`
}

var (
	AuthUrl = "https://www.linkedin.com/oauth/v2/accessToken"
)

func GenerateLinkedinAccessToken() (authresponse *AuthResponse, err error) {
	endpoint := fmt.Sprintf("?grant_type=client_credentials&client_id=%s&client_secret=%s",
		"77pky0rvu0e3x4", "xq7zLKaOXvIRbO1d")

	path := AuthUrl + endpoint
	client := http.Client{}

	request, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Println("Error making request", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Println("Error making request", err)
	}
	defer response.Body.Close()

	var res AuthResponse
	fmt.Println(response.Status)

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respByte))

	if err := json.Unmarshal(respByte, &res); err != nil {
		log.Println("Error decoding response::: ", err)
	}

	return &res, nil
}
