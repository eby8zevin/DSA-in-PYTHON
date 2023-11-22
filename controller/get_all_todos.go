package controller

import (
	"database/sql"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func GetAllTodosController(e *echo.Echo, db *sql.DB) {
	e.GET("/todos",func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM todos")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []model.TodoResponse
		for rows.Next() {
			var id int
			var title string
			var description string
			var done int
			
			err = rows.Scan(&id, &title, &description, &done)
			if err !=nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var todo model.TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1 {
				todo.Done = true
			}

			res = append(res, todo)
		}

		return ctx.JSON(http.StatusOK, res)
	})
}