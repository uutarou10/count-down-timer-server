package memorydb

import (
	"fmt"
	"github.com/uutarou10/count-down-timer-server/domain/model"
	"github.com/uutarou10/count-down-timer-server/domain/repository"
)

type MockTimerRepository struct {
	timers []model.Timer
}

func NewMockTimerRepository() repository.TimerRepository {
	repo := MockTimerRepository{
		timers: []model.Timer{},
	}

	return &repo
}

func (repository *MockTimerRepository) Create(timer model.Timer) error {
	repository.timers = append(repository.timers, timer)
	for i, timer := range repository.timers {
		fmt.Printf("%d: %+v\n", i, timer)
	}
	return nil
}

func (repository *MockTimerRepository) Update(timer model.Timer) error {
	for i, t := range repository.timers {
		if t.Id == t.Id {
			repository.timers[i] = timer
		}
	}
	return nil
}

func (repository *MockTimerRepository) Find(id string) (*model.Timer, error) {
	for _, timer := range repository.timers {
		if timer.Id == id {
			return &timer, nil
		}
	}
	return nil, fmt.Errorf("no such timer: %s", id)
}
