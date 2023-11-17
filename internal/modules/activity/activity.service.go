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

func (as *ActivityService) CreateCheckingActivityEvent(checkingActivityDto CheckingActivityDto) (ActivityEvent, error) {
	hasInteractions := checkHasInteractions(checkingActivityDto.InactivityTime, checkDuration)
	isFaceRecognized := checkIsFaceRecognized(checkingActivityDto.Photo)

	wasEmployeeActive := isFaceRecognized && hasInteractions

	var chekingActivityEventDto ActivityEventDto

	chekingActivityEventDto.CheckDuration = checkDuration
	chekingActivityEventDto.CheckTime = time.Now()
	chekingActivityEventDto.EmployeeId = checkingActivityDto.EmployeeId
	chekingActivityEventDto.EventTypeId = checkingActivityEventId
	chekingActivityEventDto.WasActive = wasEmployeeActive

	return as.activityRepository.CreateActivityEvent(chekingActivityEventDto)
}

func checkHasInteractions(inactivityTime, checkDuration int64) bool {
	hasInteractions := inactivityTime < checkDuration
	return hasInteractions
}

// TODO: implements logic
func checkIsFaceRecognized(photo os.File) bool {
    return rand.Intn(2) == 1
}

func (as *ActivityService) CreateConfirmingActivityEvent(confirmingActivityDto ConfirmingActivityDto) (ActivityEvent,error) {
	var activityEventDto ActivityEventDto
	
	activityEventDto.CheckDuration = confirmingActivityDto.CheckDuration
	activityEventDto.CheckTime = confirmingActivityDto.CheckTime
	activityEventDto.EmployeeId = confirmingActivityDto.EmployeeId
	activityEventDto.EventTypeId = confirmingActivityEventId
	activityEventDto.WasActive = true

	return as.activityRepository.CreateActivityEvent(activityEventDto)
}
