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

// checked time
const checkDuration = int64(time.Minute * 3)

func (as* ActivityService) CreateActivity(createActivityDto CreateActivityDto) (Activity, error) {
	hasInteractions := checkHasInteractions(createActivityDto.InactivityTime, checkDuration)
	isFaceRecognized := checkIsFaceRecognized(createActivityDto.Photo)

	wasEmployeeActive := isFaceRecognized && hasInteractions

	return as.activityRepository.CreateActivity(createActivityDto.EmployeeId, wasEmployeeActive, checkDuration)
}

func checkHasInteractions(inactivityTime, checkDuration int64) bool {
	hasInteractions := inactivityTime < checkDuration
	return hasInteractions
}

// TODO: implements logic
func checkIsFaceRecognized(photo os.File) bool {
    return rand.Intn(2) == 1
}

func (as *ActivityService)  AddWorkTime (employeeId int) error {
	return as.activityRepository.AddWorkTime(employeeId, checkDuration)
}

func (as *ActivityService) ConfirmActivity(id string) error {
	return as.activityRepository.ConfirmActivity(id, checkDuration)
}
