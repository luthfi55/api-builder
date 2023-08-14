package helpers

import (	
	"log"
	"github.com/jeksilaen/api-builder/modules/request/models"
	"encoding/json"
)

func ReturnSucessGetResponse(request []*models.Request) *models.SucessGetResponse {
	var requestResponses []models.RequestResponse

	for _, request := range request {
		payloadDataBytes, err := json.Marshal(request.Payload)
		if err != nil {
			log.Println("Failed to encode payload data to JSON:", err)
			return nil
		}
		responseDataBytes, err := json.Marshal(request.Response)
		if err != nil {
			log.Println("Failed to encode response data to JSON:", err)
			return nil
		}
		requestResponses = append(requestResponses, models.RequestResponse{
			ID:       request.ID,
			CollectionID:   request.CollectionID,
			Name:     request.Name,
			URL:	request.URL,
			Method:	request.Method,
			BearerToken: request.BearerToken,
			Payload:      json.RawMessage(payloadDataBytes),
			Response:     json.RawMessage(responseDataBytes),
		})
	}

	return &models.SucessGetResponse{
		Message: "Get Request sucessfully",
		Data:    requestResponses,
		Links: []models.Link{
			{
				Rel:  "get request",
				Href: "/users/v1/request",
			},
		},
	}
}

func ReturnFailedCreateRequestResponse(message string) *models.FailedResponse {
	return &models.FailedResponse{
		Error:   "Register failed",
		Message: message,
		Links: []models.Link{
			{
				Rel:  "create request",
				Href: "/users/v1/request",
			},
		},
	}
}

func convertToJSON(data interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func ReturnSucessCreateRequestResponse(createdRequest *models.Request) *models.SucessCreateResponse {
	payloadJSON, err := convertToJSON(createdRequest.Payload)
	if err != nil {
		log.Println("Failed to encode payload to JSON:", err)
		return nil
	}

	var payloadData map[string]interface{}
	err = json.Unmarshal(payloadJSON, &payloadData)
	if err != nil {
		log.Println("Failed to unmarshal JSON payload:", err)
		return nil
	}

	payloadDataBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Println("Failed to encode payload data to JSON:", err)
		return nil
	}

	responseJSON, err := convertToJSON(createdRequest.Response)
	if err != nil {
		log.Println("Failed to encode response to JSON:", err)
		return nil
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(responseJSON, &responseData)
	if err != nil {
		log.Println("Failed to unmarshal JSON response:", err)
		return nil
	}

	responseDataBytes, err := json.Marshal(responseData)
	if err != nil {
		log.Println("Failed to encode response data to JSON:", err)
		return nil
	}

	return &models.SucessCreateResponse{
		Message: "Request successfully created",
		Data: models.RequestResponse{
			ID:           createdRequest.ID,
			CollectionID: createdRequest.CollectionID,
			Name:         createdRequest.Name,
			URL:          createdRequest.URL,
			Method:       createdRequest.Method,
			BearerToken:  createdRequest.BearerToken,
			Payload:      json.RawMessage(payloadDataBytes),
			Response:     json.RawMessage(responseDataBytes),
		},		
	}
}

func ReturnSucessUpdateRequestResponse(createdRequest *models.Request) *models.SucessCreateResponse {
	payloadJSON, err := convertToJSON(createdRequest.Payload)
	if err != nil {
		log.Println("Failed to encode payload to JSON:", err)
		return nil
	}

	var payloadData map[string]interface{}
	err = json.Unmarshal(payloadJSON, &payloadData)
	if err != nil {
		log.Println("Failed to unmarshal JSON payload:", err)
		return nil
	}

	payloadDataBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Println("Failed to encode payload data to JSON:", err)
		return nil
	}

	responseJSON, err := convertToJSON(createdRequest.Response)
	if err != nil {
		log.Println("Failed to encode response to JSON:", err)
		return nil
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(responseJSON, &responseData)
	if err != nil {
		log.Println("Failed to unmarshal JSON response:", err)
		return nil
	}

	responseDataBytes, err := json.Marshal(responseData)
	if err != nil {
		log.Println("Failed to encode response data to JSON:", err)
		return nil
	}

	return &models.SucessCreateResponse{
		Message: "Request successfully updated",
		Data: models.RequestResponse{
			ID:           createdRequest.ID,
			CollectionID: createdRequest.CollectionID,
			Name:         createdRequest.Name,
			URL:          createdRequest.URL,
			Method:       createdRequest.Method,
			BearerToken:  createdRequest.BearerToken,
			Payload:      json.RawMessage(payloadDataBytes),
			Response:     json.RawMessage(responseDataBytes),
		},		
	}
}

func ReturnSucessDeleteResponse(request []*models.Request) *models.SucessDeleteResponse {
	var requestResponses []models.DeleteResponse

	for _, request := range request {		
		requestResponses = append(requestResponses, models.DeleteResponse{
			ID:       request.ID,
		})
	}

	return &models.SucessDeleteResponse{
		Message: "Delete Request sucessfully",
		Data:    requestResponses,
		Links: []models.Link{
			{
				Rel:  "delete request",
				Href: "/users/v1/request",
			},
		},
	}
}