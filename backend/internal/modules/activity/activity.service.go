package activity

import (
	"math/rand"
	"os"
	"time"
)

type ActivityService struct {
	activityRepository ActivityRepository
}

func NewActiviySerivce(activityRepository ActivityRepository) *ActivityService {
	return &ActivityService{activityRepository: activityRepository}
}

func (s* ActivityService) CreateActivity(createActivityDto CreateActivityDto) (Activity, error) {
	// checked time
	const checkDuration = int64(time.Minute * 3)
	hasInteractions := checkHasInteractions(createActivityDto.InactivityTime, checkDuration)
	isFaceRecognized := checkIsFaceRecognized(createActivityDto.Photo)

	wasEmployeeActive := isFaceRecognized && hasInteractions

	return s.activityRepository.CreateActivity(createActivityDto.EmployeeId, wasEmployeeActive, checkDuration)
}

func checkHasInteractions(inactivityTime, checkDuration int64) bool {
	hasInteractions := inactivityTime < checkDuration
	return hasInteractions
}

// TODO: implements logic
func checkIsFaceRecognized(photo os.File) bool {
    return rand.Intn(2) == 1
}

func (s *ActivityService) ConfirmActivity(id string) error {
	return s.activityRepository.ConfirmActivity(id)
}
