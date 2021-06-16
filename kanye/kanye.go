package kanye

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type kanyeTweet struct {
	Quote string `json:"quote"`
}

func (t *kanyeTweet) String() string {
	return fmt.Sprintf("Kanye says:\n\t %v", t.Quote)
}

const (
	APIUri = "https://api.kanye.rest/"
)

func GetKanyeTweet() (string, error) {
	resp, err := http.Get(APIUri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tweet kanyeTweet
	err = json.NewDecoder(resp.Body).Decode(&tweet)
	if err != nil {
		return "", nil
	}

	return tweet.String(), nil

}
