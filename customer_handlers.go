package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CustomerHandler handles customer-related requests
type CustomerHandler struct {
	db *Database
}

// NewCustomerHandler creates a new customer handler
func NewCustomerHandler(db *Database) *CustomerHandler {
	return &CustomerHandler{db: db}
}

// CreateCustomer creates a new customer
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var req CreateCustomerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	query := `INSERT INTO customers (name, email, phone, address) VALUES (?, ?, ?, ?)`
	result, err := h.db.Exec(query, req.Name, req.Email, req.Phone, req.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create customer"})
	}

	id, _ := result.LastInsertId()
	customer, err := h.GetCustomerByID(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch created customer"})
	}

	return c.JSON(http.StatusCreated, customer)
}

// GetCustomers retrieves all customers
func (h *CustomerHandler) GetCustomers(c echo.Context) error {
	query := `SELECT id, name, email, phone, address, created_at, updated_at FROM customers ORDER BY created_at DESC`
	rows, err := h.db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch customers"})
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			continue
		}
		customers = append(customers, customer)
	}

	return c.JSON(http.StatusOK, customers)
}

// GetCustomer retrieves a customer by ID
func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	customer, err := h.GetCustomerByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Customer not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch customer"})
	}

	return c.JSON(http.StatusOK, customer)
}

// GetCustomerByID helper method
func (h *CustomerHandler) GetCustomerByID(id int) (*Customer, error) {
	query := `SELECT id, name, email, phone, address, created_at, updated_at FROM customers WHERE id = ?`
	var customer Customer
	err := h.db.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// UpdateCustomer updates a customer
func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	var req CreateCustomerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	query := `UPDATE customers SET name = ?, email = ?, phone = ?, address = ? WHERE id = ?`
	result, err := h.db.Exec(query, req.Name, req.Email, req.Phone, req.Address, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update customer"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Customer not found"})
	}

	customer, err := h.GetCustomerByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch updated customer"})
	}

	return c.JSON(http.StatusOK, customer)
}

// DeleteCustomer deletes a customer
func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	query := `DELETE FROM customers WHERE id = ?`
	result, err := h.db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete customer"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Customer not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Customer deleted successfully"})
}
