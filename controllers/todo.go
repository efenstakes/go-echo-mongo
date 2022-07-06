package controllers

import (
	"go-mongo-one/models"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(c echo.Context) error {
	todo := new(models.Todo)

	if err := c.Bind(todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := mgm.Coll(todo).Create(todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}

func Get(c echo.Context) error {
	id := c.Param("id")

	todo := new(models.Todo)

	if err := mgm.Coll(todo).FindByID(id, todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func GetAll(c echo.Context) error {
	todos := []models.Todo{}

	if err := mgm.Coll(&models.Todo{}).SimpleFind(&todos, bson.M{}); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todos)
}
