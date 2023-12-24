package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/eby8zevin/golang-todos/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserController(e *echo.Echo, db *sql.DB) {
	e.POST("/auth/login", func(ctx echo.Context) error {
		var request model.LoginUserRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		row := db.QueryRow(
			"SELECT id, name, email, password FROM users WHERE email = ?", 
			request.Email,
		)
		if row.Err() != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var retrivedId int
		var retrivedName, retrivedEmail, retrivedPassword string

		err := row.Scan(&retrivedId, &retrivedName, &retrivedEmail, &retrivedPassword)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ctx.String(http.StatusUnauthorized, "email is not registered")
			}

			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		rows, err := db.Query(
			"SELECT scopes.name as scope_name FROM users LEFT JOIN user_roles ON user_roles.user_id = users.id JOIN scopes ON scopes.Id = user_roles.scope_id WHERE email = ?",
			retrivedEmail,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var scopes []string = make([]string, 0)
		for rows.Next() {
			var scope string

			err = rows.Scan(&scope)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			scopes = append(scopes, scope)
		}

		err = bcrypt.CompareHashAndPassword([]byte(retrivedPassword), []byte(request.Password))
		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}

		tokenClaim := model.AuthClaimJwt {
			UserId: retrivedId,
			UserName: retrivedName,
			UserEmail: retrivedEmail,
			UserScopes: scopes,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
		tokenStr, err := token.SignedString([]byte("TEST"))
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		response := model.LoginUserResponse {
			AccessToken: tokenStr,
		}

		return ctx.JSON(http.StatusOK, response)
	})
}