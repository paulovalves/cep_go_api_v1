package models

type ResponseModel struct {
	Data    interface{}
	Error   error
	Message string
}

func SetResponse(data interface{}, err error, message string) ResponseModel {
	if err != nil {
		return ResponseModel{Data: nil, Error: err, Message: message}
	}
	return ResponseModel{Data: data, Error: nil, Message: message}
}
