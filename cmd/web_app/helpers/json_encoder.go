package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type RegisterUserResponse struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Error    int32  `json:"error"`
}

func JsonEncoder(data map[string]string) []byte {

	jData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Json error", err)
	}

	return jData
}

func JsonDecoder(data *http.Response) RegisterUserResponse {

	var result RegisterUserResponse

	err := json.NewDecoder(data.Body).Decode(&result)
	if err != nil {
		log.Fatal("Json error", err)
	}

	return result
}
