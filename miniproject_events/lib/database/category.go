package database

import (
	"saya/config"
	"saya/models"
)

func GetCategoriesController() (interface{}, error) {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryController(CategoryID uint) (interface{}, error) {
	var category models.Category
	category.ID = CategoryID

	if err := config.DB.First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func CreateCategoryController(u models.Category) (interface{}, error) {
	err := config.DB.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func DeleteCategoryController(CategoryID int) (interface{}, error) {
	err := config.DB.Delete(&models.Category{}, CategoryID).Error

	if err != nil {
		return nil, err
	}
	return CategoryID, nil
}

func UpdateCategoryController(CategoryID uint, u models.Category) (interface{}, error) {
	category := models.Category{}
	category.ID = CategoryID
	config.DB.First(&category)

	category.Event = u.Event
	category.Category = u.Category

	err := config.DB.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
