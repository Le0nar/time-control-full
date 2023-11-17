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

const (
	// checked time
	checkDuration = int64(time.Minute * 3)
	// Event IDs
	checkingActivityEventId = 1
	confirmingActivityEventId = 2
)

func (as* ActivityService) CreateCheckingActivityEvent(checkingActivityDto CheckingActivityDto) (ActivityEvent, error) {
	hasInteractions := checkHasInteractions(checkingActivityDto.InactivityTime, checkDuration)
	isFaceRecognized := checkIsFaceRecognized(checkingActivityDto.Photo)

	wasEmployeeActive := isFaceRecognized && hasInteractions

	var chekingActivityEvent ActivityEventDto

	chekingActivityEvent.CheckDuration = checkDuration
	chekingActivityEvent.CheckTime = time.Now()
	chekingActivityEvent.EmployeeId = checkingActivityDto.EmployeeId
	chekingActivityEvent.EventTypeId = checkingActivityEventId
	chekingActivityEvent.WasActive = wasEmployeeActive

	return as.activityRepository.CreateActivityEvent(chekingActivityEvent)
}

func checkHasInteractions(inactivityTime, checkDuration int64) bool {
	hasInteractions := inactivityTime < checkDuration
	return hasInteractions
}

// TODO: implements logic
func checkIsFaceRecognized(photo os.File) bool {
    return rand.Intn(2) == 1
}
