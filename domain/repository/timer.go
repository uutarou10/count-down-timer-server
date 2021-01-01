package repository

import (
	"github.com/uutarou10/count-down-timer-server/domain/model"
)

type TimerRepository interface {
	Create(timer model.Timer) error
	Update(timer model.Timer) error
	Find(id string) (*model.Timer, error)
}
