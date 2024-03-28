package service

import (
	"github.com/le0nar/time-control-read/internal/modules/activity"
	"github.com/le0nar/time-control-read/internal/repository"
)

type Service struct {
	ActivitySerice activity.ActivityService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		ActivitySerice: *activity.NewActiviySerivce(repository.ActivityRepository),
	}
}
