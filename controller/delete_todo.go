package controller

import (
	"database/sql"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func DeleteTodoController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		result, err := db.Exec(
			"DELETE FROM todos WHERE id = ?",
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
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
			Message: "Successfully deleted",
		}

		return ctx.JSON(http.StatusOK, responseData)
	})
}