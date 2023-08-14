package usecases

import (
	"errors"
	"log"
	"strings"	
	"github.com/jeksilaen/api-builder/db"
	collectionModels"github.com/jeksilaen/api-builder/modules/collection/models"	
	requestModels"github.com/jeksilaen/api-builder/modules/request/models"	
	"gorm.io/gorm"
)

type CollectionCommandUsecase struct {
	DB *gorm.DB
}

func NewCollectionCommandUsecase() *CollectionCommandUsecase {
	return &CollectionCommandUsecase{
		DB: db.GetDB(),
	}
}

func (uc *CollectionCommandUsecase) GetCollectionsByUserID(userID string) ([]*collectionModels.Collection, error) {
	var collections []*collectionModels.Collection
	result := uc.DB.Where("user_id = ?", userID).Preload("User").Find(&collections)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("Collections not found")
		}
		return nil, result.Error
	}

	return collections, nil
}

func (uc *CollectionCommandUsecase) CreateCollection(collection *collectionModels.Collection) (*collectionModels.Collection, error) {
	
	// Create the collectiion and handle duplicate error	
	err := uc.DB.Create(collection).Error
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return nil, errors.New("User Id Not Found")
		}

		log.Println("Error creating collection:", err)
		return nil, err
	}

	return collection, nil
}


func (uc *CollectionCommandUsecase) GetCollectionByIDWithoutPreload(collectionID string) (*collectionModels.Collection, error) {
    var collection collectionModels.Collection
    result := uc.DB.Where("id = ?", collectionID).First(&collection)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("Collection not found")
        }
        return nil, result.Error
    }

    return &collection, nil
}

func (uc *CollectionCommandUsecase) UpdateCollection(collection *collectionModels.Collection) (*collectionModels.Collection, error) {
	err := uc.DB.Save(collection).Error
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (uc *CollectionCommandUsecase) DeleteCollection(collectionID string) error {
    // Check if the collection exists
    var collection collectionModels.Collection
    result := uc.DB.Where("id = ?", collectionID).First(&collection)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return errors.New("Collection not found")
        }
        return result.Error
    }

    // Delete related data from the request table
    err := uc.DB.Where("collection_id = ?", collectionID).Delete(&requestModels.Request{}).Error
    if err != nil {
        return err
    }

    // Delete the collection
    err = uc.DB.Delete(&collection).Error
    if err != nil {
        return err
    }

    return nil
}
