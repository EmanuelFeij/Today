package response

import (
	"encoding/json"
	"net/http"
)

func RequestResponse(apiUri string, destination interface{}) error {
	resp, err := http.Get(apiUri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(destination)
	if err != nil {
		return err
	}
	return nil
}
