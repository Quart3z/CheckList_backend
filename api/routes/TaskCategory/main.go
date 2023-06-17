package TaskCategory

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/quart3z/check-list/internal/database/taskCategory"
)

func SubRoutes(app *fiber.App, db *sql.DB) {

	app.Get("/taskCategories", func(c *fiber.Ctx) error {

		taskCategories := taskCategory.GetTaskCategories(db)

		json, err := json.Marshal(taskCategories)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Get("/taskCategory/:id", func(c *fiber.Ctx) error {

		taskCategory := taskCategory.GetTaskCategory(db, c.Params("id"))

		json, err := json.Marshal(taskCategory)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Post("/taskCategory", func(c *fiber.Ctx) error {

		taskCategoryJson := taskCategory.TaskCategory{}

		err := json.Unmarshal(c.Body(), &taskCategoryJson)

		if err != nil {
			panic(err.Error())
		}

		newTaskCategory, err := taskCategory.CreateTaskCategory(db, taskCategoryJson)

		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(newTaskCategory)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Put("/taskCategory", func(c *fiber.Ctx) error {

		taskCategoryJson := taskCategory.TaskCategory{}

		err := json.Unmarshal(c.Body(), &taskCategoryJson)

		if err != nil {
			panic(err.Error())
		}

		taskCategory, err := taskCategory.UpdateTaskCategory(db, taskCategoryJson)

		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(taskCategory)

		if err != nil {
			panic(err.Error())
		}

		return c.Send(json)

	})

	app.Delete("/taskCategory/:id", func(c *fiber.Ctx) error {

		taskCategory, err := taskCategory.DeleteTaskCategory(db, c.Params("id"))

		if err != nil {
			panic(err.Error())
		}

		json, err := json.Marshal(taskCategory)

		if err != nil {
			panic(err.Error())

		}

		return c.Send(json)

	})

}
