// Description: This file contains the response model for the API.
// Autor: Paulo Alves
package models

type ResponseModel struct {
	Data    interface{}
	Error   interface{}
	Message string
}

// SetResponse function to set response data for the API.
// params: data interface{}, err interface{}, message string
// returns: ResponseModel
func SetResponse(data interface{}, err interface{}, message string) ResponseModel {
	if err != nil {
		return ResponseModel{Data: nil, Error: err, Message: message}
	}
	return ResponseModel{Data: data, Error: nil, Message: message}
}
