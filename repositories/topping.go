package repositories

import (
	"BE-waysbuck-API/models"

	"gorm.io/gorm"
)

type ToppingRepository interface {
	GetToppings() ([]models.Topping, error)
	GetTopping(ID int) (models.Topping, error)
	CreateTopping(topping models.Topping) (models.Topping, error)
	UpdateTopping(topping models.Topping) (models.Topping, error)
	DeleteTopping(topping models.Topping) (models.Topping, error)
}

func RespositoryTopping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetToppings() ([]models.Topping, error) {
	var topping []models.Topping
	err := r.db.Preload("User").Find(&topping).Error

	return topping, err
}

func (r *repository) GetTopping(ID int) (models.Topping, error) {
	var topping models.Topping
	err := r.db.Preload("User").First(&topping, ID).Error

	return topping, err
}

func (r *repository) CreateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Create(&topping).Error

	return topping, err
}

func (r *repository) UpdateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Save(&topping).Error

	return topping, err
}

func (r *repository) DeleteTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Delete(&topping).Error

	return topping, err
}
