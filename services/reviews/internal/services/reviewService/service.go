package reviewService

import (
	"context"
    "reviews_service/internal/models"
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

func NewService(storage Storage, producer Producer) *Service {
    return &Service {
        storage: storage,
        producer: producer,
    }
}
// import (
//     "context"
// )

// type Service struct {
//     repo       Repository
//     producer   Producer
// }

// func NewService(repo Repository, producer Producer) *Service {
//     return &Service{repo: repo, producer: producer}
// }

// func (s *Service) CreateReview(ctx context.Context, req CreateReviewRequest) (string, error) {
//     id, err := s.repo.InsertReview(ctx, req)
//     if err != nil {
//         return "", err
//     }

//     event := RecalculationEvent{
//         ProductID: req.ProductID,
//         ReviewID:  id,
//     }

//     if err := s.producer.Send("recalculation-requests", event); err != nil {
//         return "", err
//     }

//     return id, nil
// }
