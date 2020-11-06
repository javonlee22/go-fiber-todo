package main

import (
	"github.com/gofiber/fiber"
	"github.com/javonlee22/go-fiber-todo/todo"
)

func setup() *fiber.App {
	app := fiber.New()
	app.Get("todo", todo.GetTodos)
	app.Get("todo/:ID", todo.GetTodoByID)
	app.Post("todo", todo.CreateTodo)
	app.Put("todo/:ID", todo.UpdateTodo)
	app.Delete("todo/:ID", todo.DeleteTodo)
	return app
}

func main() {
	server := setup()
	server.Listen(8000)
}
