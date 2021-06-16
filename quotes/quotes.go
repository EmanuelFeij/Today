package quotes

import (
	"math/rand"
	"time"

	"github.com/EmanuelFeij/Today/response"
)

const (
	APIUri = "https://type.fit/api/quotes"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func GetInspiratonalQuote() (string, error) {

	quotes := make([]Quote, 0, 2000)

	err := response.RequestResponse(APIUri, &quotes)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(quotes))

	return quotes[num].Text, nil
}
