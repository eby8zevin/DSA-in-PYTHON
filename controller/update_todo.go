package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func UpdateTodoController(e *echo.Echo, db *sql.DB) {
	e.PATCH("todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request model.UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		result, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ?",
			request.Title,
			request.Description,
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil{
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		if rowsAffected == 0 {
			responseData := model.DefaultResponse{
				Success: false,
				Message: "id Not found",
			}
			
			return ctx.JSON(http.StatusNotFound, responseData)
		} 

		responseData := model.DefaultResponse{
			Success: true,
			Message: "Data Updated successfully",
		}

		return ctx.JSON(http.StatusOK, responseData)
	})
}