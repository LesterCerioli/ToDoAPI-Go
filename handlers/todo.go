package handlers

import (
	models "todo-api/model"
	store "todo-api/storage"

	"github.com/gofiber/fiber/v3"
)

type TodoHandler struct {
	store *store.SQLiteStore
}

func NewTodoHandler(s *store.SQLiteStore) *TodoHandler {
	return &TodoHandler{store: s}
}

// Create godoc
// @Summary      Create a new todo
// @Description  Create a new todo item with title and optional description
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateTodoRequest  true  "Todo creation payload"
// @Success      201      {object}  map[string]interface{}    "Created todo"
// @Failure      400      {object}  map[string]string         "Invalid request"
// @Router       /todos [post]
func (h *TodoHandler) Create(c fiber.Ctx) error {
	var req models.CreateTodoRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	created, err := h.store.Create(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": created,
	})
}

// GetAll godoc
// @Summary      List all todos
// @Description  Retrieve all todo items from database
// @Tags         todos
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "List of todos"
// @Router       /todos [get]
func (h *TodoHandler) GetAll(c fiber.Ctx) error {
	todos, err := h.store.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"count": len(todos),
		"data":  todos,
	})
}

// GetByID godoc
// @Summary      Get a todo by ID
// @Description  Retrieve a specific todo item by its UUID
// @Tags         todos
// @Produce      json
// @Param        id   path      string                  true  "Todo ID"
// @Success      200  {object}  map[string]interface{}  "Todo item"
// @Failure      404  {object}  map[string]string       "Todo not found"
// @Router       /todos/{id} [get]
func (h *TodoHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID parameter is required",
		})
	}

	todo, err := h.store.GetByID(id)
	if err != nil {
		if err == store.ErrTodoNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": todo,
	})
}

// Update godoc
// @Summary      Update a todo
// @Description  Partially update a todo item (title, description, or completed status)
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id       path      string                   true  "Todo ID"
// @Param        request  body      models.UpdateTodoRequest true  "Update payload"
// @Success      200      {object}  map[string]interface{}   "Updated todo"
// @Failure      400      {object}  map[string]string        "Invalid request"
// @Failure      404      {object}  map[string]string        "Todo not found"
// @Router       /todos/{id} [patch]
func (h *TodoHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID parameter is required",
		})
	}

	var req models.UpdateTodoRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	todo, err := h.store.Update(id, &req)
	if err != nil {
		if err == store.ErrTodoNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": todo,
	})
}

// Delete godoc
// @Summary      Delete a todo
// @Description  Soft-delete a todo item (or hard-delete depending on config)
// @Tags         todos
// @Param        id   path      string  true  "Todo ID"
// @Success      204  "No Content"
// @Failure      404  {object}  map[string]string  "Todo not found"
// @Router       /todos/{id} [delete]
func (h *TodoHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID parameter is required",
		})
	}

	if err := h.store.Delete(id); err != nil {
		if err == store.ErrTodoNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
