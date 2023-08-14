package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/jeksilaen/api-builder/modules/collection/models"
	"encoding/json"
	"database/sql/driver" 
	"errors"
)

type Request struct {
	gorm.Model
	ID         string                 `gorm:"type:uuid;primaryKey"`
	CollectionID string               `gorm:"type:uuid;"`
	Name       string                 `json:"name" validate:"required"`
	URL        string                 `json:"url"`
	Method        string                 `json:"method"`
	BearerToken string					`json:"bearer_token"`
	Payload      JSONMap   `gorm:"type:json"`
    Response     JSONMap   `gorm:"type:json"`
	Collection   models.Collection `gorm:"foreignKey:CollectionID"`
}

type JSONMap map[string]interface{}

// Scan converts data from the database into the JSONMap type (map[string]interface{}).
// It unmarshals the JSON data (provided as a byte slice) into the JSONMap object.
func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshal JSONMap")
	}
	return json.Unmarshal(b, j)
}

// Value converts the JSONMap object into a JSON-encoded byte slice suitable for storage in the database.
func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type RequestResponse struct {
	ID       string `json:"id"`
	CollectionID   string `json:"collection_id"`
	Name    string `json:"name" validate:"required"`
	URL    string `json:"url"`
	Method        string                 `json:"method"`
	BearerToken string					`json:"bearer_token"`
	Payload      json.RawMessage `json:"payload"`
	Response     json.RawMessage `json:"response"`
}

type SucessCreateResponse struct {
	Message string       `json:"message"`
	Data    RequestResponse `json:"data"`	
}

type SucessGetResponse struct {
	Message string             `json:"message"`
	Data    []RequestResponse `json:"data"`
	Links   []Link             `json:"links"`
}

type DeleteResponse struct {
	ID       string `json:"id"`
}

type SucessDeleteResponse struct {
	Message string             `json:"message"`
	Data    []DeleteResponse `json:"data"`
	Links   []Link             `json:"links"`
}

type FailedResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Links   []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (request *Request) BeforeCreate(tx *gorm.DB) error {
	request.ID = uuid.New().String()
	return nil
}