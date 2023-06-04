package types

import (
	"encoding/json"
	"net/http"
)

type ResponseJSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(code int, message string) *ResponseJSON {
	return &ResponseJSON{
		Code:    code,
		Message: message,
	}
}

func NewOKResponse(message string, data interface{}) *ResponseJSON {
	var res interface{}
	byteData, ok := data.([]byte)
	if ok {
		err := json.Unmarshal(byteData, &res)
		if err != nil {
			// todo failed unmarshal the []byte result(json)
		} else {
			return &ResponseJSON{
				Code:    http.StatusOK,
				Message: message,
				Data:    res,
			}
		}
	}
	return &ResponseJSON{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
