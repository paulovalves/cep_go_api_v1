package models

type ResponseModel struct {
	data    interface{}
	err     error
	message string
}
