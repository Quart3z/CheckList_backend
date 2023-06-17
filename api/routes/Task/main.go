package Task

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/quart3z/check-list/internal/database/task"
)

func SubRoutes(app *fiber.App, db *sql.DB) {

	app.Get("/tasks", func(c *fiber.Ctx) error {

		categoryId := c.Query("category")
		if categoryId == "" {
			categoryId = "00000000-0000-0000-0000-000000000000"
		}

		tasks := task.GetTasks(db, categoryId)

		json, err := json.Marshal(tasks)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Get("/task/:id", func(c *fiber.Ctx) error {

		task := task.GetTask(db, c.Params("id"))

		json, err := json.Marshal(task)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Post("/task", func(c *fiber.Ctx) error {

		taskJson := task.Task{}
		err := json.Unmarshal(c.Body(), &taskJson)
		if err != nil {
			panic(err.Error())
		}

		task, err := task.CreateTask(db, taskJson)
		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(task)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Put("/task", func(c *fiber.Ctx) error {

		taskJson := task.Task{}
		err := json.Unmarshal(c.Body(), &taskJson)
		if err != nil {
			panic(err.Error())
		}

		updatedTask, err := task.UpdateTask(db, taskJson)

		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(updatedTask)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Put("/task/complete/:id", func(c *fiber.Ctx) error {

		task, err := task.CompleteTask(db, c.Params("id"))
		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(task)
		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Delete("/task/:id", func(c *fiber.Ctx) error {

		task, err := task.DeleteTask(db, c.Params("id"))

		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(task)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

}
