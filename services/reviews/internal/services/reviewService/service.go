package reviewService

import (
	"context"
	"errors"
	"reviews_service/internal/models"

	"github.com/google/uuid"
)


type Storage interface {
	Save(ctx context.Context, review models.ReviewInfo) error
}

type Producer interface {
	SendReviewCreated(ctx context.Context, productID int) error
}

type Service struct {
    storage Storage
    producer Producer 
    } 

func New(storage Storage, producer Producer) *Service {
    return &Service {
        storage: storage,
        producer: producer,
    }
}

func (s Service) CreateReview(ctx context.Context, productID int, rate int) (string, error) {
    if rate < 1 || rate > 5 {
		return "", errors.New("rate must be between 1 and 5")
	}

    if productID < 1 || productID > 100_000 {
        return "", errors.New("Product id must be between 1 and 100_000")
    }
	
    newID := uuid.New()
    review := models.ReviewInfo{
        ID: newID,
        ProductID: productID,
        Rate: rate,
    }

    if err := s.storage.Save(ctx, review); err != nil {
		return "", err
	}

    if err := s.producer.SendReviewCreated(ctx, productID); err != nil {
        return "", err
    }

    strID := newID.String()
    return strID, nil
}
