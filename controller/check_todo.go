package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func CheckTodoController(e *echo.Echo, db *sql.DB) {
	e.PATCH("todos/:id/check",func(ctx echo.Context) error {
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

		var request model.CheckRequest
		errReq := json.NewDecoder(ctx.Request().Body).Decode(&request)
		if errReq != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request payload")
		}

		var doneInt int
		if request.Done {
			doneInt = 1
		}

		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ? AND user_id = ?",
			doneInt,
			id,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var message string
		if request.Done {
			message = "Check successfully"
		} else {
			message = "Uncheck successfully"
		}

		responseData := model.DefaultResponse{
			Success: true,
			Message: message,
		}

		return ctx.JSON(http.StatusOK, responseData )
	})
}