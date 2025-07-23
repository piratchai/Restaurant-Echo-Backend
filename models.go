package main

import (
	"time"
)

// Restaurant represents a restaurant entity
type Restaurant struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Address     string    `json:"address" db:"address"`
	Phone       string    `json:"phone" db:"phone"`
	Email       string    `json:"email" db:"email"`
	CuisineType string    `json:"cuisine_type" db:"cuisine_type"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// MenuItem represents a menu item entity
type MenuItem struct {
	ID           int       `json:"id" db:"id"`
	RestaurantID int       `json:"restaurant_id" db:"restaurant_id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	Price        float64   `json:"price" db:"price"`
	Category     string    `json:"category" db:"category"`
	IsAvailable  bool      `json:"is_available" db:"is_available"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Customer represents a customer entity
type Customer struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Order represents an order entity
type Order struct {
	ID              int         `json:"id" db:"id"`
	CustomerID      int         `json:"customer_id" db:"customer_id"`
	RestaurantID    int         `json:"restaurant_id" db:"restaurant_id"`
	TotalAmount     float64     `json:"total_amount" db:"total_amount"`
	Status          string      `json:"status" db:"status"`
	OrderDate       time.Time   `json:"order_date" db:"order_date"`
	DeliveryAddress string      `json:"delivery_address" db:"delivery_address"`
	Notes           string      `json:"notes" db:"notes"`
	Items           []OrderItem `json:"items,omitempty"`
}

// OrderItem represents an order item entity
type OrderItem struct {
	ID         int       `json:"id" db:"id"`
	OrderID    int       `json:"order_id" db:"order_id"`
	MenuItemID int       `json:"menu_item_id" db:"menu_item_id"`
	Quantity   int       `json:"quantity" db:"quantity"`
	UnitPrice  float64   `json:"unit_price" db:"unit_price"`
	MenuItem   *MenuItem `json:"menu_item,omitempty"`
}

// CreateRestaurantRequest for creating restaurants
type CreateRestaurantRequest struct {
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	CuisineType string `json:"cuisine_type"`
}

// CreateMenuItemRequest for creating menu items
type CreateMenuItemRequest struct {
	RestaurantID int     `json:"restaurant_id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	Category     string  `json:"category"`
	IsAvailable  bool    `json:"is_available"`
}

// CreateCustomerRequest for creating customers
type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// CreateOrderRequest for creating orders
type CreateOrderRequest struct {
	CustomerID      int                      `json:"customer_id" validate:"required"`
	RestaurantID    int                      `json:"restaurant_id" validate:"required"`
	DeliveryAddress string                   `json:"delivery_address"`
	Notes           string                   `json:"notes"`
	Items           []CreateOrderItemRequest `json:"items" validate:"required,min=1"`
}

// CreateOrderItemRequest for creating order items
type CreateOrderItemRequest struct {
	MenuItemID int `json:"menu_item_id" validate:"required"`
	Quantity   int `json:"quantity" validate:"required,gt=0"`
}
