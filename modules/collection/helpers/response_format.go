package helpers

import (	
	"github.com/jeksilaen/api-builder/modules/collection/models"
)

func ReturnFailedCreateResponse(message string) *models.FailedResponse {
	return &models.FailedResponse{
		Error:   "Create failed",
		Message: message,
		Links: []models.Link{
			{
				Rel:  "create collection",
				Href: "/users/v1/collection",
			},
		},
	}
}

func ReturnSucessCreateResponse(createdCollection *models.Collection) *models.SucessCreateResponse {
	return &models.SucessCreateResponse{
		Message: "Create Collection sucessfully",
		Data: models.CollectionResponse{
			ID:       createdCollection.ID,
			UserID:       createdCollection.UserID,
			Name:    createdCollection.Name,			
		},
		Links: []models.Link{
			{
				Rel:  "Create collection",
				Href: "/users/v1/request",
			},
		},
	}
}

func ReturnSucessGetResponse(collections []*models.Collection) *models.SucessGetResponse {
	var collectionResponses []models.CollectionResponse

	for _, collection := range collections {
		collectionResponses = append(collectionResponses, models.CollectionResponse{
			ID:       collection.ID,
			UserID:   collection.UserID,
			Name:     collection.Name,
		})
	}

	return &models.SucessGetResponse{
		Message: "Get Collection sucessfully",
		Data:    collectionResponses,
		Links: []models.Link{
			{
				Rel:  "get collection",
				Href: "/users/v1/collection",
			},
		},
	}
}

func ReturnSucessDeleteResponse(message string) *models.SucessDeleteResponse {
    return &models.SucessDeleteResponse{
        Message: message,
        Links: []models.Link{
            {
                Rel:  "create collection",
                Href: "/users/v1/collection",
            },
        },
    }
}
