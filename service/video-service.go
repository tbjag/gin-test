package service

import (
	"example/web-service-gin/entity"
	"example/web-service-gin/repository"
)

type RecipeService interface {
	Save(entity.Recipe) error
	Update(entity.Recipe) error
	Delete(entity.Recipe) error
	FindAll() []entity.Recipe
}

type recipeService struct {
	repository repository.RecipeRepository
}

func New(recipeRepository repository.RecipeRepository) RecipeService {
	return &recipeService{
		repository: recipeRepository,
	}
}

func (service *recipeService) Save(recipe entity.Recipe) error {
	service.repository.Save(recipe)
	return nil
}

func (service *recipeService) Update(recipe entity.Recipe) error {
	service.repository.Update(recipe)
	return nil
}

func (service *recipeService) Delete(recipe entity.Recipe) error {
	service.repository.Delete(recipe)
	return nil
}

func (service *recipeService) FindAll() []entity.Recipe {
	return service.repository.FindAll()
}
