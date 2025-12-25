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
	s.storage = &mockStorage{
		saveFn: func(ctx context.Context, review models.ReviewInfo) error {
			assert.Equal(s.T(), 10, review.ProductID)
			assert.Equal(s.T(), 5, review.Rate)
			assert.NotZero(s.T(), review.ID)
			return nil
		},
	}

	s.producer = &mockProducer{
		sendFn: func(ctx context.Context, productID int) error {
			assert.Equal(s.T(), 10, productID)
			return nil
		},
	}

	s.service = New(s.storage, s.producer)

	id, err := s.service.CreateReview(s.ctx, 10, 5)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidRateLow() {
	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 1, 0)

	assert.Error(s.T(), err)
	assert.Empty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidRateHigh() {
	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 1, 6)

	assert.Error(s.T(), err)
	assert.Empty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidProductIDLow() {
	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 0, 3)

	assert.Error(s.T(), err)
	assert.Empty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_InvalidProductIDHigh() {
	s.service = New(nil, nil)

	id, err := s.service.CreateReview(s.ctx, 100_001, 3)

	assert.Error(s.T(), err)
	assert.Empty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_StorageError() {
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

	assert.Error(s.T(), err)
	assert.Empty(s.T(), id)
}

func (s *ReviewServiceSuite) TestCreateReview_ProducerError() {
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

	assert.Error(s.T(), err)
	assert.Empty(s.T(), id)
}

func TestReviewServiceSuite(t *testing.T) {
	suite.Run(t, new(ReviewServiceSuite))
}
