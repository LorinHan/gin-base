package rest

import "net/http"

type Rest struct {
	Status int
	//Data interface{} `json:"data"` // 加上json的话，返回时的反序列化会转为指定的字符串，这里就变成小写的d开头了
	Data interface{}
	Message string
}

func New(status int, data interface{}, message string) (int, *Rest) {
	return http.StatusOK, &Rest{status, data, message}
}

func Success(data interface{}) (int, *Rest) {
	return http.StatusOK, &Rest{http.StatusOK, data, "success"}
}

func Error(message string) (int, *Rest) {
	return http.StatusOK, &Rest{http.StatusInternalServerError, nil, message}
}
