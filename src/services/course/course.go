package course

import (
	"errors"
	"test/src/database"
	"test/src/services/lesson"
)

type Course struct {
	Id          int64           `json:"id"`
	CreatedAt   string          `json:"created_at"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Lessons     []lesson.Lesson `json:"lessons"`
}

func Get(id int64) (Course, error) {
	rows, err := database.Open("application", true).Query(`SELECT
		id, created_at, title, description
	FROM courses
	WHERE id = ?
	LIMIT 1`, id)

	if err != nil {
		return Course{}, err
	}

	courses := []Course{}

	for rows.Next() {
		course := Course{}

		rows.Scan(
			&course.Id,
			&course.CreatedAt,
			&course.Title,
			&course.Description,
		)

		courses = append(courses, course)
	}

	if len(courses) == 0 {
		return Course{}, errors.New("Course not found")
	}

	course := courses[0]

	course.Lessons, err = lesson.All(course.Id)

	if err != nil {
		return course, err
	}

	return course, nil
}

func Create(course Course) (Course, error) {
	var inserted Course
	db := database.Open("application", true)

	stmt, err := db.Exec(
		"INSERT INTO courses(title, description) VALUES(?,?);",
		course.Title,
		course.Description,
	)

	if err != nil {
		return inserted, err
	}

	lastInsertId, err := stmt.LastInsertId()

	if err != nil {
		return inserted, err
	}

	return Get(lastInsertId)
}
