package lesson

import (
	"test/src/database"
)

type Lesson struct {
	Id          int64  `json:"id"`
	CourseId    string `json:"course_id"`
	Order       int    `json:"order"`
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func All(courseId int64) ([]Lesson, error) {
	lessons := []Lesson{}

	rows, err := database.Open("application", true).Query(`SELECT
		id, created_at, title, description
	FROM lessons
	WHERE course_id = ?
	LIMIT 1`, courseId)

	if err != nil {
		return lessons, err
	}

	for rows.Next() {
		lesson := Lesson{}

		rows.Scan(
			&lesson.Id,
			&lesson.CreatedAt,
			&lesson.Title,
			&lesson.Description,
		)

		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

func Get(id int64) (Lesson, error) {
	lesson := Lesson{}

	rows, err := database.Open("application", true).Query(`SELECT
		id, created_at, order, title, description
	FROM lessons
	WHERE id = ?
	LIMIT 1`, id)

	if err != nil {
		return lesson, err
	}

	for rows.Next() {
		rows.Scan(
			&lesson.Id,
			&lesson.CreatedAt,
			&lesson.Title,
			&lesson.Description,
		)
	}

	return lesson, nil
}

func Create(lesson Lesson) (Lesson, error) {
	var inserted Lesson

	result, err := database.Open("application", true).Exec(`INSERT
	INTO lessons(course_id, order, title, description) VALUES (?,?,?,?)`,
		lesson.CourseId, lesson.Order, lesson.Title, lesson.Description)

	if err != nil {
		return inserted, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return inserted, err
	}

	return Get(lastInsertId)
}
