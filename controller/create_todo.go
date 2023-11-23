package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func CreateTodoController(e *echo.Echo, db *sql.DB) {
	e.POST("/todos",func(ctx echo.Context) error {
		var request model.CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)
		
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		responseData := model.DefaultResponse{
			Success: true,
			Message: "Data Added successfully",
		}

		return ctx.JSON(http.StatusOK, responseData)
	})
}