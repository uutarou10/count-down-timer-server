package infrastructure

import (
	"github.com/uutarou10/count-down-timer-server/domain/model"
	"github.com/uutarou10/count-down-timer-server/domain/repository"
)

type MockTimerRepository struct {
	CollectionName string
}

func NewMockTimerRepository() repository.TimerRepository {
	repo := MockTimerRepository{
		CollectionName: "timers",
	}

	return &repo
}

func (repository *MockTimerRepository) Create(timer model.Timer) error {
	return nil
}

func (repository *MockTimerRepository) Update(timer model.Timer) error {
	return nil
}

func (repository *MockTimerRepository) Find(id string) (*model.Timer, error) {
	return nil, nil
}
