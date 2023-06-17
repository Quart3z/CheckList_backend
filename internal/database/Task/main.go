package task

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func GetTasks(db *sql.DB, categoryId string) []Task {

	rows, err := db.Query(`
		SELECT task.* FROM task as task
			WHERE ($1::uuid = '00000000-0000-0000-0000-000000000000' OR task.category_id = $1)
			ORDER BY created_at DESC
	`, categoryId)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	tasks := []Task{}

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		tasks = append(tasks, task)

	}

	return tasks

}

func GetTask(db *sql.DB, id string) Task {

	rows, err := db.Query("SELECT * FROM task WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	task := Task{}

	for rows.Next() {

		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.CategoryId, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

	}

	return task

}

func CreateTask(db *sql.DB, t Task) (Task, error) {

	transaction, err := db.Begin()
	if err != nil {
		return Task{}, err
	}

	statement, err := db.Prepare("INSERT INTO task (title, description, category_id) VALUES ($1, $2, $3) RETURNING id, title, description, is_completed, category_id, created_at, updated_at")
	if err != nil {
		transaction.Rollback()
		return Task{}, err
	}

	task := Task{}
	err = statement.QueryRow(t.Title, t.Description, t.CategoryId).Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		transaction.Rollback()
		return Task{}, err
	}

	statement.Close()

	return task, nil

}

func UpdateTask(db *sql.DB, t Task) (Task, error) {

	statement, err := db.Prepare("UPDATE task SET title = $1, description = $2, is_completed = $3, category_id = $4 WHERE id = $5 RETURNING id, title, description, is_completed, category_id, created_at, updated_at")
	if err != nil {
		return Task{}, err
	}

	defer statement.Close()
	task := Task{}

	err = statement.QueryRow(t.Title, t.Description, t.IsCompleted, t.CategoryId, t.Id).Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func CompleteTask(db *sql.DB, id string) (Task, error) {

	fmt.Println("started to complete at time: ", time.Now())

	statement, err := db.Prepare("UPDATE task SET is_completed = true WHERE id = $1 RETURNING id, title, description, is_completed, category_id, created_at, updated_at")
	if err != nil {
		return Task{}, err
	}

	fmt.Println("succesffuly completed at time: ", time.Now())

	defer statement.Close()
	task := Task{}

	err = statement.QueryRow(id).Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func DeleteTask(db *sql.DB, id string) (Task, error) {

	fmt.Println("started to delete at: ", time.Now())

	statement, err := db.Prepare("DELETE FROM task WHERE id = $1 RETURNING id, title, description, is_completed, category_id, created_at, updated_at")
	if err != nil {
		return Task{}, err
	}

	fmt.Println("succesffuly deleted at time: ", time.Now())

	defer statement.Close()
	task := Task{}

	err = statement.QueryRow(id).Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

type Task struct {
	Id          string
	Title       string
	Description string
	IsCompleted bool
	CategoryId  string
	CreatedAt   string
	UpdatedAt   string
}
