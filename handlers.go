package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RestaurantHandler handles restaurant-related requests
type RestaurantHandler struct {
	db *Database
}

// NewRestaurantHandler creates a new restaurant handler
// func NewRestaurantNaja(db *Database) *RestaurantHandler {
// 	return &RestaurantHandler{db: db}
// }

// CreateRestaurant creates a new restaurant
func (h *RestaurantHandler) CreateRestaurant(c echo.Context) error {
	var req CreateRestaurantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	query := `INSERT INTO restaurants (name, address, phone, email, cuisine_type) VALUES (?, ?, ?, ?, ?)`
	result, err := h.db.Exec(query, req.Name, req.Address, req.Phone, req.Email, req.CuisineType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create restaurant"})
	}

	id, _ := result.LastInsertId()
	restaurant, err := h.GetRestaurantByID(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch created restaurant"})
	}

	return c.JSON(http.StatusCreated, restaurant)
}

// GetRestaurants retrieves all restaurants
func (h *RestaurantHandler) GetRestaurants(c echo.Context) error {
	query := `SELECT id, name, address, phone, email, cuisine_type, created_at, updated_at FROM restaurants ORDER BY created_at DESC`
	rows, err := h.db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch restaurants"})
	}
	defer rows.Close()

	var restaurants []Restaurant
	for rows.Next() {
		var r Restaurant
		err := rows.Scan(&r.ID, &r.Name, &r.Address, &r.Phone, &r.Email, &r.CuisineType, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			continue
		}
		restaurants = append(restaurants, r)
	}

	return c.JSON(http.StatusOK, restaurants)
}

// GetRestaurant retrieves a restaurant by ID
func (h *RestaurantHandler) GetRestaurant(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant ID"})
	}

	restaurant, err := h.GetRestaurantByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Restaurant not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch restaurant"})
	}

	return c.JSON(http.StatusOK, restaurant)
}

// GetRestaurantByID helper method
func (h *RestaurantHandler) GetRestaurantByID(id int) (*Restaurant, error) {
	query := `SELECT id, name, address, phone, email, cuisine_type, created_at, updated_at FROM restaurants WHERE id = ?`
	var r Restaurant
	err := h.db.QueryRow(query, id).Scan(&r.ID, &r.Name, &r.Address, &r.Phone, &r.Email, &r.CuisineType, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// UpdateRestaurant updates a restaurant
func (h *RestaurantHandler) UpdateRestaurant(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant ID"})
	}

	var req CreateRestaurantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	query := `UPDATE restaurants SET name = ?, address = ?, phone = ?, email = ?, cuisine_type = ? WHERE id = ?`
	result, err := h.db.Exec(query, req.Name, req.Address, req.Phone, req.Email, req.CuisineType, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update restaurant"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Restaurant not found"})
	}

	restaurant, err := h.GetRestaurantByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch updated restaurant"})
	}

	return c.JSON(http.StatusOK, restaurant)
}

// DeleteRestaurant deletes a restaurant
func (h *RestaurantHandler) DeleteRestaurant(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant ID"})
	}

	query := `DELETE FROM restaurants WHERE id = ?`
	result, err := h.db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete restaurant"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Restaurant not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Restaurant deleted successfully"})
}

// MenuItemHandler handles menu item-related requests
type MenuItemHandler struct {
	db *Database
}

// NewMenuItemHandler creates a new menu item handler
func NewMenuItemHandler(db *Database) *MenuItemHandler {
	return &MenuItemHandler{db: db}
}

// CreateMenuItem creates a new menu item
func (h *MenuItemHandler) CreateMenuItem(c echo.Context) error {
	var req CreateMenuItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	query := `INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := h.db.Exec(query, req.RestaurantID, req.Name, req.Description, req.Price, req.Category, req.IsAvailable)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create menu item"})
	}

	id, _ := result.LastInsertId()
	menuItem, err := h.GetMenuItemByID(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch created menu item"})
	}

	return c.JSON(http.StatusCreated, menuItem)
}

// GetMenuItems retrieves menu items (optionally filtered by restaurant)
func (h *MenuItemHandler) GetMenuItems(c echo.Context) error {
	restaurantID := c.QueryParam("restaurant_id")

	var query string
	var args []interface{}

	if restaurantID != "" {
		query = `SELECT id, restaurant_id, name, description, price, category, is_available, created_at, updated_at FROM menu_items WHERE restaurant_id = ? ORDER BY category, name`
		args = append(args, restaurantID)
	} else {
		query = `SELECT id, restaurant_id, name, description, price, category, is_available, created_at, updated_at FROM menu_items ORDER BY restaurant_id, category, name`
	}

	rows, err := h.db.Query(query, args...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch menu items"})
	}
	defer rows.Close()

	var menuItems []MenuItem
	for rows.Next() {
		var m MenuItem
		err := rows.Scan(&m.ID, &m.RestaurantID, &m.Name, &m.Description, &m.Price, &m.Category, &m.IsAvailable, &m.CreatedAt, &m.UpdatedAt)
		if err != nil {
			continue
		}
		menuItems = append(menuItems, m)
	}

	return c.JSON(http.StatusOK, menuItems)
}

// GetMenuItem retrieves a menu item by ID
func (h *MenuItemHandler) GetMenuItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu item ID"})
	}

	menuItem, err := h.GetMenuItemByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Menu item not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch menu item"})
	}

	return c.JSON(http.StatusOK, menuItem)
}

// GetMenuItemByID helper method
func (h *MenuItemHandler) GetMenuItemByID(id int) (*MenuItem, error) {
	query := `SELECT id, restaurant_id, name, description, price, category, is_available, created_at, updated_at FROM menu_items WHERE id = ?`
	var m MenuItem
	err := h.db.QueryRow(query, id).Scan(&m.ID, &m.RestaurantID, &m.Name, &m.Description, &m.Price, &m.Category, &m.IsAvailable, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateMenuItem updates a menu item
func (h *MenuItemHandler) UpdateMenuItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu item ID"})
	}

	var req CreateMenuItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	query := `UPDATE menu_items SET restaurant_id = ?, name = ?, description = ?, price = ?, category = ?, is_available = ? WHERE id = ?`
	result, err := h.db.Exec(query, req.RestaurantID, req.Name, req.Description, req.Price, req.Category, req.IsAvailable, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update menu item"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Menu item not found"})
	}

	menuItem, err := h.GetMenuItemByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch updated menu item"})
	}

	return c.JSON(http.StatusOK, menuItem)
}

// DeleteMenuItem deletes a menu item
func (h *MenuItemHandler) DeleteMenuItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu item ID"})
	}

	query := `DELETE FROM menu_items WHERE id = ?`
	result, err := h.db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete menu item"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Menu item not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Menu item deleted successfully"})
}
