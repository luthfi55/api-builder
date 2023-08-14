package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	"github.com/jeksilaen/api-builder/middlewares"
	"github.com/jeksilaen/api-builder/modules/request/helpers"	
	"github.com/jeksilaen/api-builder/modules/request/models"	
	"github.com/jeksilaen/api-builder/modules/request/usecases"
)

func InitRequestHttpHandler(router *gin.Engine) {	
	router.GET("/users/v1/request/:request_id", middlewares.VerifyToken, GetRequestById)
	router.GET("/users/v1/request_by_collection/:collection_id", middlewares.VerifyToken, GetRequestByCollection)
	router.POST("/users/v1/request", middlewares.VerifyToken,CreateRequest)
	router.PUT("/users/v1/request/:request_id", middlewares.VerifyToken,UpdateRequest)
	router.DELETE("/users/v1/request/:request_id", middlewares.VerifyToken,DeleteRequest)
}

func GetRequestById(ctx *gin.Context) {
	requestUsecase := usecases.NewRequestCommandUsecase()

	requestID := ctx.Param("request_id")

	if requestID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request ID is required"})
		return
	}

	// Get the request data from usecase
	request, err := requestUsecase.GetRequestByRequestID(requestID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	// Return the request data as response
	ctx.JSON(http.StatusCreated, helpers.ReturnSucessGetResponse([]*models.Request{request}))
}

func GetRequestByCollection(ctx *gin.Context) {
	requestUsecase := usecases.NewRequestCommandUsecase()

	collectionID := ctx.Param("collection_id")

	if collectionID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Collection ID is required"})
		return
	}

	// Get the request data from usecase
	request, err := requestUsecase.GetRequestByCollectionID(collectionID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	// Return the request data as response
	ctx.JSON(http.StatusCreated, helpers.ReturnSucessGetResponse(request))
}


func CreateRequest(ctx *gin.Context) {
	requestUsecase := usecases.NewRequestCommandUsecase()
	// validate := validator.New()

	// Decode the request JSON data into User object
	var req models.Request
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateRequestResponse(err.Error()))
		return
	}

	// Validate the request JSON data
	// err := validate.Struct(req)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateRequestResponse(err.Error()))
	// 	return
	// }
	
	// Create the user
	createdCollection, _ := requestUsecase.CreateRequest(&req)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateRequestResponse(err.Error()))
	// 	return
	// }

	ctx.JSON(http.StatusCreated, helpers.ReturnSucessCreateRequestResponse(createdCollection))
	
}

func UpdateRequest(ctx *gin.Context) {
    requestUsecase := usecases.NewRequestCommandUsecase()
    // validate := validator.New()

    // Get collection ID from path parameter
    requestID := ctx.Param("request_id")

    // Validate collection ID (optional, based on your requirements)
    if requestID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request ID is required"})
        return
    }

    // Decode the request JSON data into Collection object
    var req models.Request
    if err := ctx.BindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateRequestResponse(err.Error()))
        return
    }

    // Validate the request JSON data
    // err := validate.StructPartial(req, "Name")
    // if err != nil {
    //     ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
    //     return
    // }

    // Get the existing collection data from usecase without preloading the User field
    existingRequest, err := requestUsecase.GetRequestByIDWithoutPreload(requestID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
        return
    }

    // Update the Name field of the existing collection
    existingRequest.Name = req.Name
	existingRequest.URL = req.URL
	existingRequest.Method = req.Method
	existingRequest.Payload = req.Payload
	existingRequest.Response = req.Response

    // Save the updated request
    updatedRequest, err := requestUsecase.UpdateRequest(existingRequest)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateRequestResponse(err.Error()))
        return
    }

    ctx.JSON(http.StatusOK, helpers.ReturnSucessUpdateRequestResponse(updatedRequest))
}

func DeleteRequest(ctx *gin.Context) {
	requestUsecase := usecases.NewRequestCommandUsecase()

	requestID := ctx.Param("request_id")

	if requestID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request ID is required"})
		return
	}

	// Get the request data from usecase
	request, err := requestUsecase.DeleteRequestByRequestID(requestID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	// Return the request data as response
	ctx.JSON(http.StatusCreated, helpers.ReturnSucessDeleteResponse([]*models.Request{request}))
}