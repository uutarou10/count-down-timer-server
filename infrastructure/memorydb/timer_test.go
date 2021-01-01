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

	t.Run("対象のエンティティがUpdateされること", func(t *testing.T) {
		repository := NewMockTimerRepository()
		targetUuid := uuid.New().String()

		_ = repository.Create(model.Timer{
			Id:        targetUuid,
			Title:     "test",
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
			return
		}

		if timer, err := repository.Find(targetUuid); err != nil {
			t.Error(err.Error())
		} else if timer.Title != "updated title" {
			t.Errorf("Expected: updated title, Actual: %s", timer.Title)
		}
	})

	t.Run("対象ではないエンティティが更新されないこと", func(t *testing.T) {
		repository := NewMockTimerRepository()
		targetUuid := uuid.New().String()
		nonTargetUuid := uuid.New().String()

		_ = repository.Create(model.Timer{
			Id:        targetUuid,
			Title:     "test",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		})

		_ = repository.Create(model.Timer{
			Id:        nonTargetUuid,
			Title:     "nontarget",
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
			return
		}

		if timer, err := repository.Find(targetUuid); err != nil {
			t.Error(err.Error())
			return
		} else if timer.Title != "updated title" {
			t.Errorf("Expected: updated title, Actual: %s", timer.Title)
			return
		}

		if timer, err := repository.Find(nonTargetUuid); err != nil {
			t.Errorf(err.Error())
		} else if timer.Title != "nontarget" {
			t.Errorf("Expected: nontarget, Actual: %s", timer.Title)
			return
		}
	})

	t.Run("作成したエンティティをFindで取得できること", func(t *testing.T) {
		repository := NewMockTimerRepository()
		targetUuid := uuid.New().String()

		_ = repository.Create(model.Timer{
			Id:        targetUuid,
			Title:     "test",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		})

		_ = repository.Create(model.Timer{
			Id:        uuid.New().String(),
			Title:     "test2",
			DueDate:   time.Now(),
			CreatedAt: time.Now(),
		})

		if timer, err := repository.Find(targetUuid); err != nil {
			t.Error(err.Error())
		} else if timer.Id != targetUuid {
			t.Errorf("Expected: %s, Actual: %s", targetUuid, timer.Id)
		}
	})
}
