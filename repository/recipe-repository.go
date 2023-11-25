package repository

import (
	"example/web-service-gin/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type RecipeRepository interface {
	Save(recipe entity.Recipe)
	Update(recipe entity.Recipe)
	Delete(recipe entity.Recipe)
	FindAll() []entity.Recipe
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewRecipeRepository() RecipeRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Recipe{}, &entity.Person{}, &entity.Ingredient{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(recipe entity.Recipe) {
	db.connection.Create(&recipe)
}

func (db *database) Update(recipe entity.Recipe) {
	db.connection.Save(&recipe)
}

func (db *database) Delete(recipe entity.Recipe) {
	db.connection.Delete(&recipe)
}

func (db *database) FindAll() []entity.Recipe {
	var recipes []entity.Recipe
	db.connection.Set("gorm:auto_preload", true).Find(&recipes)
	return recipes
}
