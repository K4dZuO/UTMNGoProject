package reviewService

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"reviews_service/internal/models"
)

type mockStorage struct {
	saveFn func(ctx context.Context, review models.ReviewInfo) error
}

func (m *mockStorage) Save(ctx context.Context, review models.ReviewInfo) error {
	return m.saveFn(ctx, review)
}

type mockProducer struct {
	sendFn func(ctx context.Context, productID int) error
}

func (m *mockProducer) SendReviewCreated(ctx context.Context, productID int) error {
	return m.sendFn(ctx, productID)
}

func TestService_CreateReview(t *testing.T) {
	rand.Seed(42)

	type args struct {
		productID int
		rate      int
	}

	tests := []struct {
		name      string
		args      args
		storage   *mockStorage
		producer  *mockProducer
		wantErr   bool
		wantID    bool
	}{
		{
			name: "success",
			args: args{
				productID: 10,
				rate:      5,
			},
			storage: &mockStorage{
				saveFn: func(ctx context.Context, review models.ReviewInfo) error {
					if review.ProductID != 10 || review.Rate != 5 {
						return errors.New("invalid review data")
					}
					return nil
				},
			},
			producer: &mockProducer{
				sendFn: func(ctx context.Context, productID int) error {
					if productID != 10 {
						return errors.New("invalid product id")
					}
					return nil
				},
			},
			wantErr: false,
			wantID:  true,
		},
		{
			name: "invalid rate low",
			args: args{
				productID: 1,
				rate:      0,
			},
			storage:  nil,
			producer: nil,
			wantErr:  true,
			wantID:   false,
		},
		{
			name: "invalid rate high",
			args: args{
				productID: 1,
				rate:      6,
			},
			storage:  nil,
			producer: nil,
			wantErr:  true,
			wantID:   false,
		},
		{
			name: "invalid product id low",
			args: args{
				productID: 0,
				rate:      3,
			},
			storage:  nil,
			producer: nil,
			wantErr:  true,
			wantID:   false,
		},
		{
			name: "invalid product id high",
			args: args{
				productID: 100_001,
				rate:      3,
			},
			storage:  nil,
			producer: nil,
			wantErr:  true,
			wantID:   false,
		},
		{
			name: "storage error",
			args: args{
				productID: 5,
				rate:      4,
			},
			storage: &mockStorage{
				saveFn: func(ctx context.Context, review models.ReviewInfo) error {
					return errors.New("storage failed")
				},
			},
			producer: &mockProducer{
				sendFn: func(ctx context.Context, productID int) error {
					return nil
				},
			},
			wantErr: true,
			wantID:  false,
		},
		{
			name: "producer error",
			args: args{
				productID: 7,
				rate:      2,
			},
			storage: &mockStorage{
				saveFn: func(ctx context.Context, review models.ReviewInfo) error {
					return nil
				},
			},
			producer: &mockProducer{
				sendFn: func(ctx context.Context, productID int) error {
					return errors.New("kafka failed")
				},
			},
			wantErr: true,
			wantID:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			var storage Storage
			var producer Producer

			if tt.storage != nil {
				storage = tt.storage
			}
			if tt.producer != nil {
				producer = tt.producer
			}

			svc := New(storage, producer)

			id, err := svc.CreateReview(ctx, tt.args.productID, tt.args.rate)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantID && id == "" {
				t.Fatalf("expected id, got empty string")
			}
		})
	}
}
