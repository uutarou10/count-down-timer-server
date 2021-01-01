package memorydb

import (
	"github.com/google/uuid"
	"github.com/uutarou10/count-down-timer-server/domain/model"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	t.Run("正常にCreateができる", func(t *testing.T) {
		repository := NewMockTimerRepository()

		if err := repository.Create(model.Timer{
			Id:        uuid.New().String(),
			Title:     "test",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		}); err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("対象のエンティティだけがUpdateされること", func(t *testing.T) {
		repository := NewMockTimerRepository()
		targetUuid := uuid.New().String()
		otherUuid := uuid.New().String()

		_ = repository.Create(model.Timer{
			Id:        targetUuid,
			Title:     "test",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		})

		_ = repository.Create(model.Timer{
			Id:        otherUuid,
			Title:     "test2",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		})

		if err := repository.Update(model.Timer{
			Id:        targetUuid,
			Title:     "updated title",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		}); err != nil {
			t.Errorf(err.Error())
		}

		if timer, _ := repository.Find(targetUuid); timer.Title != "updated title" {
			t.Errorf("Expected: updated title, Actual: %s", timer.Title)
		}

		if timer, err := repository.Find(otherUuid); err != nil {
			t.Errorf(err.Error())
		} else if timer.Title == "updated title" {
			t.Errorf("hoge")
		}
	})
}
