package controller

import (
	"example/web-service-gin/entity"
	"example/web-service-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecipeController interface {
	FindAll() []entity.Recipe
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.RecipeService
}

func New(service service.RecipeService) RecipeController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Recipe {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var recipe entity.Recipe
	err := ctx.ShouldBindJSON(&recipe)
	if err != nil {
		return err
	}
	c.service.Save(recipe)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var recipe entity.Recipe
	err := ctx.ShouldBindJSON(&recipe)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	recipe.ID = id

	c.service.Update(recipe)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var recipe entity.Recipe
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	recipe.ID = id
	c.service.Delete(recipe)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	recipes := c.service.FindAll()
	data := gin.H{
		"title":  "Recipe Page",
		"recipes": recipes,
	}
	ctx.HTML(http.StatusOK, "allrecipe.html", data)
}
