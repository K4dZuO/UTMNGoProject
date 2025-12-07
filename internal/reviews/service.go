package reviews

import (
    "context"
    "go_back/internal/kafka"
)

type Service struct {
    repo       Repository
    producer   kafka.Producer
}

func NewService(repo Repository, producer kafka.Producer) *Service {
    return &Service{repo: repo, producer: producer}
}

func (s *Service) CreateReview(ctx context.Context, req CreateReviewRequest) (string, error) {
    id, err := s.repo.InsertReview(ctx, req)
    if err != nil {
        return "", err
    }

    event := kafka.RecalculationEvent{
        ProductID: req.ProductID,
        ReviewID:  id,
    }

    if err := s.producer.Send("recalculation-requests", event); err != nil {
        return "", err
    }

    return id, nil
}
