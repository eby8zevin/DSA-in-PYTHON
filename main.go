package main

import (
	"github.com/eby8zevin/golang-todos/controller"
	"github.com/eby8zevin/golang-todos/database"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()
	
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	controller.GetAllTodosController(e, db)

	e.Start(":8080")
}