package usecase

import (
	"github.com/uutarou10/count-down-timer-server/domain/model"
	"github.com/uutarou10/count-down-timer-server/domain/repository"
)

type TimerUsecase interface {
	Add(timer model.Timer) error
	Get(id string) (*model.Timer, error)
}

type timerUsecase struct {
	repository repository.TimerRepository
}

func (usecase *timerUsecase) Add(timer model.Timer) error {
	return usecase.repository.Create(timer)
}

func (usecase *timerUsecase) Update(timer model.Timer) error {
	return usecase.repository.Update(timer)
}

func (usecase *timerUsecase) Get(id string) (*model.Timer, error) {
	if timer, err := usecase.repository.Find(id); err != nil {
		return nil, err
	} else {
		return timer, nil
	}
}
