package activity

type ActivityService struct {
	activityRepository ActivityRepository
}

func NewActiviySerivce(activityRepository ActivityRepository) *ActivityService {
	return &ActivityService{activityRepository: activityRepository}
}

func (as *ActivityService) GetEmployeeMonthActivity(employeeId, year, month int) ([]DayActivityDto, error) {
	return as.activityRepository.GetEmployeeMonthActivity(employeeId, year, month)
}

func (as *ActivityService) AddWorkTime(dto AddingWorkingTimeDto) error {
	return as.activityRepository.AddWorkTime(dto)
}
