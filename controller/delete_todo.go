package controller

import (
	"database/sql"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func DeleteTodoController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		user := ctx.Get("USER").(model.AuthClaimJwt)

		permissionFound := false
		for _, scope := range user.UserScopes {
			if scope == "todos:delete" {
				permissionFound = true
				break
			}
		}
		if !permissionFound {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		id := ctx.Param("id")

		result, err := db.Exec(
			"DELETE FROM todos WHERE id = ? AND user_id = ?",
			id,
			user.UserId,
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