package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:id`
	Completed bool   `json:completed`
	Text      string `json:text`
}

var todos []Todo // slice of todos

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	// check if text is null
	if c.BodyParser(todo) != nil {
		fmt.Println("/api/v1/todos(POST) -> says 'cannot parse JSON'")
		return c.Status(400).JSON(fiber.Map{"err": "cannot parse JSON"})
	}

	// check if text is empty
	if todo.Text == "" {
		fmt.Println("/api/v1/todos(POST) -> says 'text is required'")
		return c.Status(400).JSON(fiber.Map{"err": "text is required"})
	}

	// create a new TODO
	todo.ID = len(todos) + 1
	todos = append(todos, *todo)
	fmt.Println("/api/v1/todos(POST) -> says 'todo created'")
	return c.Status(200).JSON(fiber.Map{"message": "todo created"})
}

func getTodo(c *fiber.Ctx) error {
	//check if todos is empty
	if len(todos) == 0 {
		fmt.Println("/api/v1/todos(GET) -> says 'no todos found'")
		return c.Status(404).JSON(fiber.Map{"err": "no todos found"})
	}

	// display todos
	fmt.Println("/api/v1/todos(GET) -> display todos")
	return c.Status(200).JSON(todos)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id") //get id from url

	// check if id is empty
	if id == "" {
		fmt.Println("/api/v1/todos/:id(PATCH) -> says 'id is required'")
		return c.Status(400).JSON(fiber.Map{"err": "id is required"})
	}

	// update todo
	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i].Completed = !todos[i].Completed // if true make if false, if false make it true
			fmt.Println("/api/v1/todos/:id(PATCH) -> says 'successfully updated'")
			return c.Status(200).JSON(fiber.Map{"message": "successfully updated"})
		}
	}

	// todo not found
	fmt.Println("/api/v1/todos/:id(PATCH) -> says 'todo not found'")
	return c.Status(400).JSON(fiber.Map{"err": "todo not pound"})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id") //get if from url

	// check if id is empty
	if id == "" {
		fmt.Println("/api/v1/todos/:id(DELETE) -> says 'id is required'")
		return c.Status(400).JSON(fiber.Map{"err": "id is required"})
	}

	// delete todo
	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("/api/v1/todos/:id(DELETE) -> says 'successfully deleted'")
			return c.Status(200).JSON(fiber.Map{"message": "successfully deleted"})
		}
	}

	// todo not found
	fmt.Println("/api/v1/todos/:id(DELETE) -> says 'todo not found'")
	return c.Status(400).JSON(fiber.Map{"err": "todo not pound"})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("/ -> says 'hello world'")
		return c.Status(200).JSON(fiber.Map{"message": "hello world"})
	})

	app.Post("/api/v1/todos", createTodo)
	app.Get("/api/v1/todos", getTodo)
	app.Patch("/api/v1/todos/:id", updateTodo)
	app.Delete("/api/v1/todos/:id", deleteTodo)

	app.Listen(":3000") // server is listening on port 3000
}
