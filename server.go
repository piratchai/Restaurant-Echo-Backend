package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRestaurantNaja(db *Database) *RestaurantHandler {
	return &RestaurantHandler{db: db}
}

func main() {
	// Initialize database
	db, err := NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize handlers
	restaurantHandler := NewRestaurantNaja(db)
	menuItemHandler := NewMenuItemHandler(db)
	customerHandler := NewCustomerHandler(db)
	orderHandler := NewOrderHandler(db)

	

	// Root endpoint
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Restaurant API",
			"version": "1.0.0",
		})
	})

	// API v1 routes
	v1 := e.Group("/api/v1")

	// Restaurant routes
	restaurants := v1.Group("/restaurants")
	restaurants.POST("", restaurantHandler.CreateRestaurant)
	restaurants.GET("", restaurantHandler.GetRestaurants)
	restaurants.GET("/:id", restaurantHandler.GetRestaurant)
	restaurants.PUT("/:id", restaurantHandler.UpdateRestaurant)
	restaurants.DELETE("/:id", restaurantHandler.DeleteRestaurant)

	// Menu item routes
	menuItems := v1.Group("/menu-items")
	menuItems.POST("", menuItemHandler.CreateMenuItem)
	menuItems.GET("", menuItemHandler.GetMenuItems) // Supports ?restaurant_id=X filter
	menuItems.GET("/:id", menuItemHandler.GetMenuItem)
	menuItems.PUT("/:id", menuItemHandler.UpdateMenuItem)
	menuItems.DELETE("/:id", menuItemHandler.DeleteMenuItem)

	// Customer routes
	customers := v1.Group("/customers")
	customers.POST("", customerHandler.CreateCustomer)
	customers.GET("", customerHandler.GetCustomers)
	customers.GET("/:id", customerHandler.GetCustomer)
	customers.PUT("/:id", customerHandler.UpdateCustomer)
	customers.DELETE("/:id", customerHandler.DeleteCustomer)

	// Order routes
	orders := v1.Group("/orders")
	orders.POST("", orderHandler.CreateOrder)
	orders.GET("", orderHandler.GetOrders) // Supports ?customer_id=X and ?restaurant_id=X filters
	orders.GET("/:id", orderHandler.GetOrder)
	orders.PATCH("/:id/status", orderHandler.UpdateOrderStatus)
	orders.DELETE("/:id", orderHandler.DeleteOrder)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Start server
	log.Println("Starting server on :3644")
	e.Logger.Fatal(e.Start(":3644"))
}
