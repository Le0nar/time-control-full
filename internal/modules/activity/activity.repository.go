package activity

import (
	"github.com/jmoiron/sqlx"
)

type ActivityRepository struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

const dayActivityTable = "activity"

func (ar *ActivityRepository) AddWorkTime(employeeId int, checkDuration int64) error {
	// year, month, day := time.Now().Date()

	// // 1) get year, or create year if it doesnt exist
	// var yearActivity YearActivity

	// // Check is year activity exist
	// query := fmt.Sprintf("SELECT * FROM %s WHERE employee_id=$1 and year=$2", yearTable)
	// err := ar.db.Get(&yearActivity, query, employeeId, year)

	// // Create new year activity if it doesnt exist
	// if err != nil {
	// 	query := fmt.Sprintf(
	// 		"INSERT INTO %s (employee_id, year) values ($1, $2) RETURNING  id, employee_id, year",
	// 		yearTable,
	// 	)
	
	// 	row := ar.db.QueryRow(query, employeeId, year)
	
	// 	err := row.Scan(&yearActivity.Id, &yearActivity.EmployeeId, &yearActivity.Year)

	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // 2) get month or create month if it doesnt exist
	// var monthActivity MonthActivity

	// // Check is month activity exist
	// query = fmt.Sprintf("SELECT * FROM %s WHERE year_id=$1 ", monthTable)
	// err = ar.db.Get(&monthActivity, query, yearActivity.Id)

	// // Create new month activity if it doesnt exist
	// if err != nil {
	// 	query := fmt.Sprintf(
	// 		"INSERT INTO %s (year_id, month) values ($1, $2) RETURNING  id, year_id, month",
	// 		monthTable,
	// 	)
	
	// 	row := ar.db.QueryRow(query, yearActivity.Id, month)
	
	// 	err := row.Scan(&monthActivity.Id, &monthActivity.YearId, &monthActivity.Month)

	// 	if err != nil {
	// 		return err
	// 	}
	// }

	
	// // 3) get day or create day if it doesnt exist, and set ac
	// var dayActivity DayActivity

	// // Check is day activity exist
	// query = fmt.Sprintf("SELECT * FROM %s WHERE month_id=$1 ", dayTable)
	// err = ar.db.Get(&dayActivity, query, monthActivity.Id)

	// // Create new day activity if it doesnt exist
	// if err != nil {
	// 	query := fmt.Sprintf(
	// 		"INSERT INTO %s (month_id, day, activity_time) values ($1, $2, $3)",
	// 		dayTable,
	// 	)
	
	// 	row := ar.db.QueryRow(query, monthActivity.Id, day, checkDuration)
	
	// 	err := row.Scan()

	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// fmt.Printf("DayActivity: %+v\n", dayActivity)

	// // Update day activity time if day activity already exists
	// if err == nil {
	// 	dayActivity.ActivityTime += checkDuration

	// 	query := fmt.Sprintf("UPDATE %s SET activity_time=$1 WHERE id=$2", dayTable)

	// 	_, err := ar.db.Exec(query, dayActivity.ActivityTime, dayActivity.Id)
	
	// 	return err
	// }
	
	return nil
}

func (ar *ActivityRepository) GetEmployeeMonthActivity(employeeId, year, month int) (MonthActivityDto, error) {
	var yearActivity YearActivity

	// get and return MonthActivityDto
}

// var newsList []news.News
// query := fmt.Sprintf("SELECT * FROM %s", newsTable)
// err := r.db.Select(&newsList, query)

// return newsList, err

// var newsItem news.News
// query := fmt.Sprintf("SELECT * FROM %s where id = %d", newsTable, newsId)
// err := r.db.Get(&newsItem, query)

// return newsItem, err