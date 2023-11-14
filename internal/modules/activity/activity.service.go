package activity

type ActivityService struct {
	activityRepository ActivityRepository
}

func NewActiviySerivce(activityRepository ActivityRepository) *ActivityService {
	return &ActivityService{activityRepository: activityRepository}
}

// TODO: mb send curr
func (as *ActivityService) AddWorkTime(employeeId int, checkDuration int64, activity_date string) error {
	return as.activityRepository.AddWorkTime(employeeId, checkDuration)
}

func (as *ActivityService) GetEmployeeMonthActivity(employeeId, year, month int) (MonthActivityDto, error) {
	return as.activityRepository.GetEmployeeMonthActivity(employeeId, year, month)
}
