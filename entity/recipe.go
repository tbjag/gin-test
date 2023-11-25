package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int    `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256);UNIQUE"`
}

type Recipe struct {
	ID             uint64       `gorm:"primary_key;auto_increment" json:"id"`
	Title          string       `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Description    string       `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	ImageURL       string       `json:"imageurl" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author         Person       `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID       uint64       `json:"-"`
	IngredientList []Ingredient `json:"ingredientlist" binding:"required" gorm:"foreignkey:RecipeID"`
	CreatedAt      time.Time    `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time    `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Ingredient struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:"name" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Weight   string `json:"weight" binding:"max=200" gorm:"type:varchar(200)"` //all will be in metric - convert to imperial
	Note     string `json:"note" binding:"max=200" gorm:"type:varchar(200)"`
	RecipeID uint64 `json:"-"`
}
