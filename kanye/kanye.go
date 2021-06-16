package kanye

import (
	"fmt"

	"github.com/EmanuelFeij/Today/response"
)

type KanyeTweet struct {
	Quote string `json:"quote"`
}

func (t *KanyeTweet) String() string {
	return fmt.Sprintf("Kanye says:\n\t %v", t.Quote)
}

const (
	APIUri = "https://api.kanye.rest/"
)

func GetKanyeTweet() (string, error) {

	var tweet KanyeTweet
	err := response.RequestResponse(APIUri, &tweet)
	if err != nil {
		return "", nil
	}

	return tweet.String(), nil

}
