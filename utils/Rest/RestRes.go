package Rest

import "net/http"

type Rest struct {
	Status int
	//Data interface{} `json:"data"` // 加上json的话，返回时的反序列化会转为指定的字符串，这里就变成小写的d开头了
	Data interface{}
	Message string
}

func Res(status int, data interface{}, message string) *Rest{
	return &Rest{status, data, message}
}

func Success(data interface{}) *Rest {
	return &Rest{http.StatusOK, data, "success"}
}

func Error(message string) *Rest {
	return &Rest{http.StatusInternalServerError, nil, message}
}
