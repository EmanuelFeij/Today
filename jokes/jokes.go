package jokes

const (
	APIUri = "https://v2.jokeapi.dev/joke/Any"
)

type jokeResponse struct {
	Error      bool        `json:"error"`
	Category   string      `json:"category"`
	TypeOfJoke string      `json:"type"`
	Setup      string      `json:"setup"`
	Delivery   string      `json:"delivery"`
	Flags      interface{} `json:"flags"`
	Id         int         `json:"id"`
	Safe       bool        `json:"safe"`
	Lang       string      `json:"lang"`
}
