package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// OrderHandler handles order-related requests
type OrderHandler struct {
	db *Database
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(db *Database) *OrderHandler {
	return &OrderHandler{db: db}
}

// CreateOrder creates a new order with items
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Start transaction
	tx, err := h.db.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start transaction"})
	}
	defer tx.Rollback()

	// Calculate total amount
	var totalAmount float64
	for _, item := range req.Items {
		// Get menu item price
		var price float64
		err := tx.QueryRow("SELECT price FROM menu_items WHERE id = ?", item.MenuItemID).Scan(&price)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu item"})
		}
		totalAmount += price * float64(item.Quantity)
	}

	// Create order
	query := `INSERT INTO orders (customer_id, restaurant_id, total_amount, delivery_address, notes) VALUES (?, ?, ?, ?, ?)`
	result, err := tx.Exec(query, req.CustomerID, req.RestaurantID, totalAmount, req.DeliveryAddress, req.Notes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}

	orderID, _ := result.LastInsertId()

	// Create order items
	for _, item := range req.Items {
		// Get menu item price again for order item
		var price float64
		err := tx.QueryRow("SELECT price FROM menu_items WHERE id = ?", item.MenuItemID).Scan(&price)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu item"})
		}

		itemQuery := `INSERT INTO order_items (order_id, menu_item_id, quantity, unit_price) VALUES (?, ?, ?, ?)`
		_, err = tx.Exec(itemQuery, orderID, item.MenuItemID, item.Quantity, price)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order items"})
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
	}

	// Fetch the created order with items
	order, err := h.GetOrderByID(int(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch created order"})
	}

	return c.JSON(http.StatusCreated, order)
}

// GetOrders retrieves orders (optionally filtered by customer or restaurant)
func (h *OrderHandler) GetOrders(c echo.Context) error {
	customerID := c.QueryParam("customer_id")
	restaurantID := c.QueryParam("restaurant_id")

	var query string
	var args []interface{}

	baseQuery := `SELECT id, customer_id, restaurant_id, total_amount, status, order_date, delivery_address, notes FROM orders`

	if customerID != "" && restaurantID != "" {
		query = baseQuery + ` WHERE customer_id = ? AND restaurant_id = ? ORDER BY order_date DESC`
		args = append(args, customerID, restaurantID)
	} else if customerID != "" {
		query = baseQuery + ` WHERE customer_id = ? ORDER BY order_date DESC`
		args = append(args, customerID)
	} else if restaurantID != "" {
		query = baseQuery + ` WHERE restaurant_id = ? ORDER BY order_date DESC`
		args = append(args, restaurantID)
	} else {
		query = baseQuery + ` ORDER BY order_date DESC`
	}

	rows, err := h.db.Query(query, args...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.CustomerID, &order.RestaurantID, &order.TotalAmount, &order.Status, &order.OrderDate, &order.DeliveryAddress, &order.Notes)
		if err != nil {
			continue
		}

		// Get order items for each order
		items, _ := h.getOrderItems(order.ID)
		order.Items = items

		orders = append(orders, order)
	}

	return c.JSON(http.StatusOK, orders)
}

// GetOrder retrieves an order by ID
func (h *OrderHandler) GetOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	order, err := h.GetOrderByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch order"})
	}

	return c.JSON(http.StatusOK, order)
}

// GetOrderByID helper method
func (h *OrderHandler) GetOrderByID(id int) (*Order, error) {
	query := `SELECT id, customer_id, restaurant_id, total_amount, status, order_date, delivery_address, notes FROM orders WHERE id = ?`
	var order Order
	err := h.db.QueryRow(query, id).Scan(&order.ID, &order.CustomerID, &order.RestaurantID, &order.TotalAmount, &order.Status, &order.OrderDate, &order.DeliveryAddress, &order.Notes)
	if err != nil {
		return nil, err
	}

	// Get order items
	items, err := h.getOrderItems(order.ID)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return &order, nil
}

// getOrderItems helper method
func (h *OrderHandler) getOrderItems(orderID int) ([]OrderItem, error) {
	query := `
		SELECT oi.id, oi.order_id, oi.menu_item_id, oi.quantity, oi.unit_price,
		       mi.name, mi.description, mi.category
		FROM order_items oi
		JOIN menu_items mi ON oi.menu_item_id = mi.id
		WHERE oi.order_id = ?
	`

	rows, err := h.db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []OrderItem
	for rows.Next() {
		var item OrderItem
		var menuItem MenuItem

		err := rows.Scan(&item.ID, &item.OrderID, &item.MenuItemID, &item.Quantity, &item.UnitPrice,
			&menuItem.Name, &menuItem.Description, &menuItem.Category)
		if err != nil {
			continue
		}

		menuItem.ID = item.MenuItemID
		menuItem.Price = item.UnitPrice
		item.MenuItem = &menuItem

		items = append(items, item)
	}

	return items, nil
}

// UpdateOrderStatus updates order status
func (h *OrderHandler) UpdateOrderStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Validate status
	validStatuses := []string{"pending", "confirmed", "preparing", "ready", "delivered", "cancelled"}
	isValid := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValid = true
			break
		}
	}
	if !isValid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid status"})
	}

	query := `UPDATE orders SET status = ? WHERE id = ?`
	result, err := h.db.Exec(query, req.Status, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order status"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}

	order, err := h.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch updated order"})
	}

	return c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order
func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	query := `DELETE FROM orders WHERE id = ?`
	result, err := h.db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete order"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Order deleted successfully"})
}
