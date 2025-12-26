package reviewService

import (
	"context"
	"errors"
	"testing"
	"time"

	"reviews_service/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

type ReviewServiceSuite struct {
	suite.Suite

	ctx      context.Context
	cancel   context.CancelFunc
	storage  *mockStorage
	producer *mockProducer
	service  *Service
}

func (s *ReviewServiceSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), time.Second)
}

func (s *ReviewServiceSuite) TearDownTest() {
	s.cancel()
}

func (s *ReviewServiceSuite) TestCreateReview_Success() {
	wantProductID := 10
	wantRate := 5

	s.storage = &mockStorage{
		saveFn: func(ctx context.Context, review models.ReviewInfo) error {
			assert.Equal(s.T(), wantProductID, review.ProductID)
			assert.Equal(s.T(), wantRate, review.Rate)
			assert.NotZero(s.T(), review.ID)
			return nil
		},
	}

	s.producer = &mockProducer{
		sendFn: func(ctx context.Context, productID int) error {
			assert.Equal(s.T(), wantProductID, productID)
			return nil
		},
	}

	s.service = New(s.storage, s.producer)

	id, err := s.service.CreateReview(s.ctx, wantProductID, wantRate)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidRateLow() {
	wantErr := true
	wantID := ""

	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 1, 0)

	assert.Equal(s.T(), wantErr, err != nil)
	assert.Equal(s.T(), wantID, id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidRateHigh() {
	wantErr := true
	wantID := ""

	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 1, 6)

	assert.Equal(s.T(), wantErr, err != nil)
	assert.Equal(s.T(), wantID, id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidProductIDLow() {
	wantErr := true
	wantID := ""

	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 0, 3)

	assert.Equal(s.T(), wantErr, err != nil)
	assert.Equal(s.T(), wantID, id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidProductIDHigh() {
	wantErr := true
	wantID := ""

	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 100_001, 3)

	assert.Equal(s.T(), wantErr, err != nil)
	assert.Equal(s.T(), wantID, id)
}

func (s *ReviewServiceSuite) TestCreateReview_StorageError() {
	wantErr := true
	wantID := ""

	s.storage = &mockStorage{
		saveFn: func(ctx context.Context, review models.ReviewInfo) error {
			return errors.New("storage error")
		},
	}

	s.producer = &mockProducer{
		sendFn: func(ctx context.Context, productID int) error {
			return nil
		},
	}

	s.service = New(s.storage, s.producer)

	id, err := s.service.CreateReview(s.ctx, 5, 4)

	assert.Equal(s.T(), wantErr, err != nil)
	assert.Equal(s.T(), wantID, id)
}

func (s *ReviewServiceSuite) TestCreateReview_ProducerError() {
	wantErr := true
	wantID := ""

	s.storage = &mockStorage{
		saveFn: func(ctx context.Context, review models.ReviewInfo) error {
			return nil
		},
	}

	s.producer = &mockProducer{
		sendFn: func(ctx context.Context, productID int) error {
			return errors.New("producer error")
		},
	}

	s.service = New(s.storage, s.producer)

	id, err := s.service.CreateReview(s.ctx, 7, 2)

	assert.Equal(s.T(), wantErr, err != nil)
	assert.Equal(s.T(), wantID, id)
}

func TestReviewServiceSuite(t *testing.T) {
	suite.Run(t, new(ReviewServiceSuite))
}
