package helpers

import (
	"net/http"
)

func CallGatewayPOST(url string, token string, payload map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("POST", url, JsonEncoder(payload))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	return client.Do(req)
}
