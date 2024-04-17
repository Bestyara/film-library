package repository

import (
	"context"
	"errors"
	mock_database "film-library/internal/db/mocks"
	"film-library/internal/model"
	"film-library/internal/service"
	"github.com/golang/mock/gomock"
	"testing"
)

var (
	ctx               = context.Background()
	id                = int64(1)
	ErrObjectNotFound = errors.New("not found")
	film              = model.Film{}
)

type mockstruct struct {
	ctrl   *gomock.Controller
	mockDB *mock_database.MockDBops
	repo   service.Repos
}

func mockinit(t *testing.T) mockstruct {
	ctrl := gomock.NewController(t)
	mockDB := mock_database.NewMockDBops(ctrl)
	repo := NewRepository(mockDB)
	return mockstruct{
		ctrl:   ctrl,
		repo:   repo,
		mockDB: mockDB,
	}

}
