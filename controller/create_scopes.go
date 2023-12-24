package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/labstack/echo"
)

func CreateScopeController(e *echo.Echo, db *sql.DB) {
	e.POST("/scopes", func(ctx echo.Context) error {
		var request model.CreateScopeRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		row := db.QueryRow(
			"SELECT id FROM scopes WHERE name = ?",
			request.Name,
		)
		if row.Err() != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var retrivedId int
		err := row.Scan(&retrivedId)
		if err == nil {
			return ctx.String(http.StatusBadRequest, "duplicate scope found")
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		_, err = db.Exec(
			"INSERT INTO scopes (name) VALUES (?)",
			request.Name,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}