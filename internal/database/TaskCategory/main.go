package taskCategory

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetTaskCategories(db *sql.DB) []TaskCategory {

	rows, err := db.Query("SELECT * FROM taskcategory ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	taskCategories := []TaskCategory{}

	for rows.Next() {
		taskCategory := TaskCategory{}

		err := rows.Scan(&taskCategory.Id, &taskCategory.Title, &taskCategory.IsDefault, &taskCategory.CreatedAt, &taskCategory.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		taskCategories = append(taskCategories, taskCategory)

	}

	return taskCategories

}

func GetTaskCategory(db *sql.DB, id string) TaskCategory {

	rows, err := db.Query("SELECT * FROM taskcategory WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	taskCategory := TaskCategory{}

	for rows.Next() {

		err := rows.Scan(&taskCategory.Id, &taskCategory.Title, &taskCategory.IsDefault, &taskCategory.CreatedAt, &taskCategory.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

	}

	return taskCategory

}

func CreateTaskCategory(db *sql.DB, t TaskCategory) (TaskCategory, error) {

	statement, err := db.Prepare("INSERT INTO taskcategory (title, is_default) VALUES ($1, $2) RETURNING id, title, is_default, created_at, updated_at")
	if err != nil {
		return TaskCategory{}, err
	}

	defer statement.Close()
	taskCategory := TaskCategory{}

	err = statement.QueryRow(t.Title, t.IsDefault).Scan(&taskCategory.Id, &taskCategory.Title, &taskCategory.IsDefault, &taskCategory.CreatedAt, &taskCategory.UpdatedAt)
	if err != nil {
		return TaskCategory{}, err
	}

	return taskCategory, nil

}

func UpdateTaskCategory(db *sql.DB, t TaskCategory) (TaskCategory, error) {

	statement, err := db.Prepare("UPDATE taskcategory SET title = $1, is_default = $2 WHERE id = $3 RETURNING id, title, is_default, created_at, updated_at")
	if err != nil {
		return TaskCategory{}, err
	}

	defer statement.Close()
	taskCategory := TaskCategory{}

	err = statement.QueryRow(t.Title, t.IsDefault, t.Id).Scan(&taskCategory.Id, &taskCategory.Title, &taskCategory.IsDefault, &taskCategory.CreatedAt, &taskCategory.UpdatedAt)
	if err != nil {
		return TaskCategory{}, err
	}

	return taskCategory, nil
}

func DeleteTaskCategory(db *sql.DB, id string) (TaskCategory, error) {

	statement, err := db.Prepare("DELETE FROM taskcategory WHERE id = $1 RETURNING id, title, is_default, created_at, updated_at")
	if err != nil {
		return TaskCategory{}, err
	}

	defer statement.Close()
	taskCategory := TaskCategory{}

	err = statement.QueryRow(id).Scan(&taskCategory.Id, &taskCategory.Title, &taskCategory.IsDefault, &taskCategory.CreatedAt, &taskCategory.UpdatedAt)
	if err != nil {
		return TaskCategory{}, err
	}

	return taskCategory, nil
}

type TaskCategory struct {
	Id        string
	Title     string
	IsDefault bool
	CreatedAt string
	UpdatedAt string
}
