package activity

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
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
	isFaceRecognized, err := checkIsFaceRecognized(&checkingActivityDto.Photo)

	if err != nil {
		// TODO: refactor: return pointer to ActivityEvent
		return ActivityEvent{}, err
	}

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

// TODO: check it 
func checkIsFaceRecognized(file *os.File) (bool, error) {
	endpoint := "http://localhost:8003/recognise"

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "exmaple.jpg")
	if err != nil {
	  return false, err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
	  return false, err
	}
	request, err := http.NewRequest("POST", endpoint, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
	  return false, err
	}
	defer response.Body.Close()

	// TODO: refactor it
	var bodyString string
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}

	if bodyString == "false" {
		return false, nil
	}
	
	return true, nil
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
