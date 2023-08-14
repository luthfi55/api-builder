package handlers

import (
	"net/http"
	// "log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jeksilaen/api-builder/middlewares"
	"github.com/jeksilaen/api-builder/modules/collection/helpers"	
	"github.com/jeksilaen/api-builder/modules/collection/models"	
	"github.com/jeksilaen/api-builder/modules/collection/usecases"
)

func InitCollectionHttpHandler(router *gin.Engine) {	
	router.GET("/users/v1/collection/:user_id", middlewares.VerifyToken, GetCollectionByUserID)
	router.POST("/users/v1/collection", middlewares.VerifyToken,CreateCollection)
	router.PUT("/users/v1/collection/:id", middlewares.VerifyToken,UpdateCollection)
	router.DELETE("/users/v1/collection/:id", middlewares.VerifyToken,DeleteCollection)
}


func GetCollectionByUserID(ctx *gin.Context) {
	collectionUsecase := usecases.NewCollectionCommandUsecase()

	// Get user_id from path parameter
	userID := ctx.Param("user_id")

	// Validate user_id (optional, based on your requirements)
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Get the collection data from usecase
	collections, err := collectionUsecase.GetCollectionsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Collections not found"})
		return
	}

	// Return the collection data as response
	// ctx.JSON(http.StatusOK, collections)
	ctx.JSON(http.StatusCreated, helpers.ReturnSucessGetResponse(collections))
}

func CreateCollection(ctx *gin.Context) {
	collectionUsecase := usecases.NewCollectionCommandUsecase()
	validate := validator.New()

	// Decode the request JSON data into Collection object
	var req models.Collection
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
		return
	}

	// Validate the request JSON data
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
		return
	}

	// Create the collection
	createdCollection, err := collectionUsecase.CreateCollection(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.ReturnSucessCreateResponse(createdCollection))
}

func UpdateCollection(ctx *gin.Context) {
    collectionUsecase := usecases.NewCollectionCommandUsecase()
    validate := validator.New()

    // Get collection ID from path parameter
    collectionID := ctx.Param("id")

    // Validate collection ID (optional, based on your requirements)
    if collectionID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Collection ID is required"})
        return
    }

    // Decode the request JSON data into Collection object
    var req models.Collection
    if err := ctx.BindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
        return
    }

    // Validate the request JSON data
    err := validate.StructPartial(req, "Name")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
        return
    }

    // Get the existing collection data from usecase without preloading the User field
    existingCollection, err := collectionUsecase.GetCollectionByIDWithoutPreload(collectionID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
        return
    }

    // Update the Name field of the existing collection
    existingCollection.Name = req.Name

    // Save the updated collection
    updatedCollection, err := collectionUsecase.UpdateCollection(existingCollection)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedCreateResponse(err.Error()))
        return
    }

    ctx.JSON(http.StatusOK, helpers.ReturnSucessCreateResponse(updatedCollection))
}

func DeleteCollection(ctx *gin.Context) {
    collectionUsecase := usecases.NewCollectionCommandUsecase()

    // Get collection ID from path parameter
    collectionID := ctx.Param("id")

    // Validate collection ID (optional, based on your requirements)
    if collectionID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Collection ID is required"})
        return
    }

    // Delete the collection
    err := collectionUsecase.DeleteCollection(collectionID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
        return
    }

    // Return the success response
    ctx.JSON(http.StatusOK, helpers.ReturnSucessDeleteResponse("Deleted Collection Successfully"))
}