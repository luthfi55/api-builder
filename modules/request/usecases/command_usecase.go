package usecases

import (
	"errors"
	"log"
	"strings"	
	"encoding/json"
	"net/http"
	"bytes"
	"strconv"
	// "encoding/json"
	"github.com/jeksilaen/api-builder/db"
	"github.com/jeksilaen/api-builder/modules/request/models"	
	"gorm.io/gorm"
)

type RequestCommandUsecase struct {
	DB *gorm.DB
}

func NewRequestCommandUsecase() *RequestCommandUsecase {
	return &RequestCommandUsecase{
		DB: db.GetDB(),
	}
}

func (uc *RequestCommandUsecase) GetRequestByRequestID(requestID string) (*models.Request, error) {
	var request models.Request
	result := uc.DB.Where("id = ?", requestID).Preload("Collection").First(&request)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("Request not found")
		}
		return nil, result.Error
	}

	return &request, nil
}

func (uc *RequestCommandUsecase) GetRequestByCollectionID(collectionID string) ([]*models.Request, error) {
	var request []*models.Request
	result := uc.DB.Where("collection_id = ?", collectionID).Preload("Collection").Find(&request)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("Request not found")
		}
		return nil, result.Error
	}

	return request, nil
}

func (uc *RequestCommandUsecase) CreateRequest(request *models.Request) (*models.Request, error) {
	
	if request.Method == "GET" {
		req, err := http.NewRequest("GET", request.URL, nil)		

		// Set the authorization header
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)

		// Make the HTTP request
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			var responseData models.JSONMap
			err := json.NewDecoder(response.Body).Decode(&responseData)
			if err != nil {
				return nil, err
			}
			request.Response = responseData
		} else {
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
			request.Response = errorResponse
		}
	} else if request.Method == "POST"{		
		payloadJSON, _ := json.Marshal(request.Payload)
			
		req, err := http.NewRequest("POST", request.URL, bytes.NewBuffer(payloadJSON))
		if err != nil {
			return nil, err
		}

		// Set the authorization header
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)
		req.Header.Set("Content-Type", "application/json")

		response, err := http.DefaultClient.Do(req)
		if err != nil {			
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Failed to fetch URL: " + err.Error()
				
			request.Response = errorResponse
		} else {
			if response.StatusCode == http.StatusOK {				
				var responseData models.JSONMap
				err := json.NewDecoder(response.Body).Decode(&responseData)
				if err != nil {					
					errorResponse := make(models.JSONMap)
					errorResponse["error"] = "Failed to parse JSON: " + err.Error()
						
					request.Response = errorResponse
				} else {					
					request.Response = responseData
				}
			} else {				
				var responseData models.JSONMap
				json.NewDecoder(response.Body).Decode(&responseData)
				request.Response = responseData
				if request.Response == nil {
					errorResponse := make(models.JSONMap)
					errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
					request.Response = errorResponse
				}
			}
		}

	}else if request.Method == "PUT" || request.Method == "PATCH" {		
		payloadJSON, _ := json.Marshal(request.Payload)		
	
		req, _ := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(payloadJSON))
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)
		req.Header.Set("Content-Type", "application/json")
		
		client := &http.Client{}
		response, err := client.Do(req)
	
		if err != nil {			
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Failed to fetch URL: " + err.Error()
				
			request.Response = errorResponse
		} else {
			var responseData models.JSONMap
			
			json.NewDecoder(response.Body).Decode(&responseData)			
			request.Response = responseData
			if request.Response == nil {
				errorResponse := make(models.JSONMap)
				errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
				request.Response = errorResponse
			}
		}	

	} else if request.Method == "DELETE" {		
		req, err := http.NewRequest("DELETE", request.URL, nil)
		
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)
		client := http.DefaultClient
		response, err := client.Do(req)

		if err != nil {			
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Failed to fetch URL: " + err.Error()
				
			request.Response = errorResponse
		} else {
			var responseData models.JSONMap
			
			json.NewDecoder(response.Body).Decode(&responseData)			
			request.Response = responseData
			if request.Response == nil {
				errorResponse := make(models.JSONMap)
				errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
				request.Response = errorResponse
			}
		}		
	}  else {
		return nil, errors.New("Invalid request method for updating, use GET, POST, PUT, or DELETE")
	}
	
	err := uc.DB.Create(request).Error
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return nil, errors.New("User Id Not Found")
		}

		log.Println("Error creating user:", err)
		return nil, err
	}

	return request, nil
}

func (uc *RequestCommandUsecase) GetRequestByIDWithoutPreload(requestID string) (*models.Request, error) {
    var request models.Request
    result := uc.DB.Where("id = ?", requestID).First(&request)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("Request not found")
        }
        return nil, result.Error
    }

    return &request, nil
}

func (uc *RequestCommandUsecase) UpdateRequest(request *models.Request) (*models.Request, error) {
	
	if request.Method == "GET" {
		req, err := http.NewRequest("GET", request.URL, nil)		

		// Set the authorization header
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)

		// Make the HTTP request
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			var responseData models.JSONMap
			err := json.NewDecoder(response.Body).Decode(&responseData)
			if err != nil {
				return nil, err
			}
			request.Response = responseData
		} else {
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
			request.Response = errorResponse
		}
	} else if request.Method == "POST"{		
		payloadJSON, _ := json.Marshal(request.Payload)
			
		req, err := http.NewRequest("POST", request.URL, bytes.NewBuffer(payloadJSON))
		if err != nil {
			return nil, err
		}

		// Set the authorization header
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)
		req.Header.Set("Content-Type", "application/json")

		response, err := http.DefaultClient.Do(req)
		if err != nil {			
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Failed to fetch URL: " + err.Error()
				
			request.Response = errorResponse
		} else {
			if response.StatusCode == http.StatusOK {				
				var responseData models.JSONMap
				err := json.NewDecoder(response.Body).Decode(&responseData)
				if err != nil {					
					errorResponse := make(models.JSONMap)
					errorResponse["error"] = "Failed to parse JSON: " + err.Error()
						
					request.Response = errorResponse
				} else {					
					request.Response = responseData
				}
			} else {				
				var responseData models.JSONMap
				json.NewDecoder(response.Body).Decode(&responseData)
				request.Response = responseData
				if request.Response == nil {
					errorResponse := make(models.JSONMap)
					errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
					request.Response = errorResponse
				}
			}
		}

	}else if request.Method == "PUT" || request.Method == "PATCH" {		
		payloadJSON, _ := json.Marshal(request.Payload)		
	
		req, _ := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(payloadJSON))
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)
		req.Header.Set("Content-Type", "application/json")
		
		client := &http.Client{}
		response, err := client.Do(req)
	
		if err != nil {			
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Failed to fetch URL: " + err.Error()
				
			request.Response = errorResponse
		} else {
			var responseData models.JSONMap
			
			json.NewDecoder(response.Body).Decode(&responseData)			
			request.Response = responseData
			if request.Response == nil {
				errorResponse := make(models.JSONMap)
				errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
				request.Response = errorResponse
			}
		}	

	} else if request.Method == "DELETE" {		
		req, err := http.NewRequest("DELETE", request.URL, nil)
		
		req.Header.Set("Authorization", "Bearer "+request.BearerToken)
		client := http.DefaultClient
		response, err := client.Do(req)

		if err != nil {			
			errorResponse := make(models.JSONMap)
			errorResponse["error"] = "Failed to fetch URL: " + err.Error()
				
			request.Response = errorResponse
		} else {
			var responseData models.JSONMap
			
			json.NewDecoder(response.Body).Decode(&responseData)			
			request.Response = responseData
			if request.Response == nil {
				errorResponse := make(models.JSONMap)
				errorResponse["error"] = "Invalid status code: " + strconv.Itoa(response.StatusCode)
				request.Response = errorResponse
			}
		}		
	}  else {
		return nil, errors.New("Invalid request method for updating, use GET, POST, PUT, or DELETE")
	}
		
	err := uc.DB.Save(request).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (uc *RequestCommandUsecase) DeleteRequestByRequestID(requestID string) (*models.Request, error) {
	var request models.Request
	result := uc.DB.Where("id = ?", requestID).Preload("Collection").First(&request)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("Request not found")
		}
		return nil, result.Error
	}

	// Delete the request from the database
	if err := uc.DB.Delete(&request).Error; err != nil {
		return nil, err
	}

	return &request, nil
}