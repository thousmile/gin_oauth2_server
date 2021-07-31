package domain

import (
	"gin_oauth2_server/exception"
	"net/http"
)

func result(status int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

func Ok1(message string, data interface{}) map[string]interface{} {
	return result(http.StatusOK, message, data)
}

func Ok2(data interface{}) map[string]interface{} {
	return result(http.StatusOK, "ok", data)
}

func Error1(status int, message string) map[string]interface{} {
	return result(status, message, nil)
}

func Error2(err *exception.OAuth2Error) map[string]interface{} {
	return result(err.Status, err.Message, nil)
}

func Error3(err *exception.OAuth2Error, message string) map[string]interface{} {
	return result(err.Status, message, nil)
}
