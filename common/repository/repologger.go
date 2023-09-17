package repository

import (
	"context"
	"time"

	"github.com/correooke/MyMeal/common/model"
	"github.com/sirupsen/logrus"
)

type LoggingRepository[T model.Entity] struct {
	repo CommonRepository[T]
}

func NewLoggingRepository[T model.Entity](r CommonRepository[T]) *LoggingRepository[T] {
	return &LoggingRepository[T]{repo: r}
}

var log = logrus.New()

func startLogging(actionName string, fields logrus.Fields) time.Time {
	if fields != nil {
		log.WithFields(fields).Infof("Starting %s", actionName)
	} else {
		log.Infof("Starting %s", actionName)
	}
	return time.Now()
}

func finishLogging(actionName string, startTime time.Time, err error) {
	elapsedTime := time.Since(startTime)
	if err != nil {
		log.WithField("duration", elapsedTime).Errorf("Error in %s: %v", actionName, err)
	} else {
		log.WithField("duration", elapsedTime).Infof("Finished %s", actionName)
	}
}

func (lr *LoggingRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	startTime := startLogging("GetAll", nil)
	result, err := lr.repo.GetAll(ctx)
	finishLogging("GetAll", startTime, err)
	return result, err
}

func (lr *LoggingRepository[T]) GetByFilter(ctx context.Context, filter map[string]interface{}) ([]T, error) {
	startTime := startLogging("GetByFilter", logrus.Fields{"filter": filter})
	result, err := lr.repo.GetByFilter(ctx, filter)
	finishLogging("GetByFilter", startTime, err)
	return result, err
}

func (lr *LoggingRepository[T]) GetByID(ctx context.Context, id string) (T, error) {
	startTime := startLogging("GetByID", logrus.Fields{"id": id})
	result, err := lr.repo.GetByID(ctx, id)
	finishLogging("GetByID", startTime, err)
	return result, err
}

func (lr *LoggingRepository[T]) Create(ctx context.Context, entity T) error {
	startTime := startLogging("Create", logrus.Fields{"entity": entity})
	err := lr.repo.Create(ctx, entity)
	finishLogging("Create", startTime, err)
	return err
}

func (lr *LoggingRepository[T]) Update(ctx context.Context, id string, entity T) error {
	startTime := startLogging("Update", logrus.Fields{"id": id, "entity": entity})
	err := lr.repo.Update(ctx, id, entity)
	finishLogging("Update", startTime, err)
	return err
}

func (lr *LoggingRepository[T]) Delete(ctx context.Context, id string) error {
	startTime := startLogging("Delete", logrus.Fields{"id": id})
	err := lr.repo.Delete(ctx, id)
	finishLogging("Delete", startTime, err)
	return err
}
