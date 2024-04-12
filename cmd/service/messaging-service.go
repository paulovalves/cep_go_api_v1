package service

import (
	"encoding/json"
	"log"

	entity "models/entity"
)

type MessageBody struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

func ProcessMessage(message []byte) entity.ResponseModel {
	var body MessageBody
	err := json.Unmarshal(message, &body)
	if err != nil {
		log.Printf("Failed to unmarshal message: %s", message)
		return entity.ResponseModel{Error: err}
	}

	switch body.Type {
	case "category":
		var category entity.Category
		err := json.Unmarshal(message, &category)
		if err != nil {
			log.Printf("Failed to unmarshal category message: %s", message)
			return entity.ResponseModel{Error: err}
		}
		res := CreateCategory(category)
		return res
	case "image":
		var image entity.Image
		err := json.Unmarshal(message, &image)
		if err != nil {
			log.Printf("Failed to unmarshal image message: %s", message)
			return entity.ResponseModel{Error: err}
		}
		// Process image...
		// res := ProcessImage(image)
		// return res
	default:
		log.Printf("Unknown message type: %s", body.Type)
		return entity.ResponseModel{Error: err}
	}
	return entity.ResponseModel{}
}
