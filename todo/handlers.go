package todo

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
)

var todosMap = make(map[uint64]*Todo)

// CreateTodo Create Todo handler
func CreateTodo(c *fiber.Ctx) {
	t := NewTodo()
	if err := c.BodyParser(t); err != nil {
		fmt.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	todosMap[t.ID] = t
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Todo created successfully",
		"todo":    t,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating todo",
		})
		return
	}
}

// GetTodos  Get Todos handler
func GetTodos(c *fiber.Ctx) {
	todosList := []*Todo{}
	for _, v := range todosMap {
		todosList = append(todosList, v)
	}
	c.JSON(todosList)
}

// GetTodoByID fetches a Todo with a matching ID
func GetTodoByID(c *fiber.Ctx) {
	id, err := strconv.ParseUint(c.Params("ID"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	if todo, ok := todosMap[id]; ok {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Fetch successful",
			"todo":    todo,
		})
		return
	}
}

// DeleteTodo Delete Todo handler
func DeleteTodo(c *fiber.Ctx) {
	k, err := strconv.ParseUint(c.Params("ID"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	if _, ok := todosMap[k]; ok {
		delete(todosMap, k)
	}
}

// UpdateTodo Upcate Todo handler
func UpdateTodo(c *fiber.Ctx) {
	k, err := strconv.ParseUint(c.Params("ID"), 10, 64)
	if err != nil {

	}
	req := &UpdateTodoRequest{}
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	if v, ok := todosMap[k]; ok {
		v.Update(req.Task, req.IsCompleted)
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Todo created successfully",
			"todo":    v})
		return
	}
}
