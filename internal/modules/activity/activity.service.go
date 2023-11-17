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

//  handle first type of event
func (as* ActivityService) CreateActivityEvent(checkingActivityDto CheckingActivityDto) (bool, error) {
	hasInteractions := checkHasInteractions(checkingActivityDto.InactivityTime, checkDuration)
	isFaceRecognized := checkIsFaceRecognized(checkingActivityDto.Photo)

	wasEmployeeActive := isFaceRecognized && hasInteractions

	var chekingActivityEvent ActivityEvent

	chekingActivityEvent.CheckDuration = checkDuration
	chekingActivityEvent.CheckTime = time.Now()
	chekingActivityEvent.EmployeeId = checkingActivityDto.EmployeeId
	chekingActivityEvent.EventTypeId = checkingActivityEventId
	chekingActivityEvent.WasActive = wasEmployeeActive

	err := as.activityRepository.CreateActivityEvent(chekingActivityEvent)

	return chekingActivityEvent.WasActive, err
}

func checkHasInteractions(inactivityTime, checkDuration int64) bool {
	hasInteractions := inactivityTime < checkDuration
	return hasInteractions
}

// TODO: implements logic
func checkIsFaceRecognized(photo os.File) bool {
    return rand.Intn(2) == 1
}

