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
		user := ctx.Get("USER").(model.AuthClaimJwt)

		permissionFound := false
		for _, scope := range user.UserScopes {
			if scope == "todos:update" {
				permissionFound = true
				break
			}
		}
		if !permissionFound {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		id := ctx.Param("id")

		var request model.UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		result, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ? AND user_id = ?",
			request.Title,
			request.Description,
			id,
			user.UserId,
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