package utils

import "fmt"

type TelarError struct {
	Error ErrorCodeMessage `json:"error"`
}

type ErrorCodeMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func MarshalError(code string, message string) []byte {
	return []byte(fmt.Sprintf(`{
		"error": {
			"message": "%s",
			"code": "%s" 
		}
	}`, message, code))
}

func Error(code string, message string) TelarError {
	return TelarError{
		Error: ErrorCodeMessage{
			Code:    code,
			Message: message,
		},
	}
}
