package repositories

import (
	"BE-waysbuck-API/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	GetProductCart(ID int) (models.Product, error)
	GetToppingCart(ID []int) ([]models.Topping, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}

func (r *repository) GetProductCart(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repository) GetToppingCart(ToppingID []int) ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Debug().Find(&toppings, ToppingID).Error

	return toppings, err
}
