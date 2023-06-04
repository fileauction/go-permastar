package model

import (
	"encoding/json"
)

type UserInfo struct {
	AccountAddress string `json:"accountAddress"`
}

type RegisterRequest struct {
	AccountAddress string `json:"accountAddress"`
}

func ReadRegisterRequest(data []byte) (*RegisterRequest, error) {
	var registerRequest *RegisterRequest
	if err := json.Unmarshal(data, &registerRequest); err != nil {
		return nil, err
	}

	return registerRequest, nil
}
