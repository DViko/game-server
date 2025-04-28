package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RegisterUserResponse struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Error    int32  `json:"error"`
}

func JsonEncoder(data map[string]string) *bytes.Buffer {

	jData, _ := json.Marshal(data)

	return bytes.NewBuffer(jData)
}

func JsonDecoder(data *http.Response) RegisterUserResponse {

	var result RegisterUserResponse

	json.NewDecoder(data.Body).Decode(&result)

	return result
}
