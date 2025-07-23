# Frontend Development Guide - Restaurant API

This guide helps frontend developers create HTML, JavaScript, and AJAX applications to connect with the Restaurant Backend API.

## Table of Contents
- [API Overview](#api-overview)
- [Authentication & CORS](#authentication--cors)
- [Basic HTML Setup](#basic-html-setup)
- [JavaScript API Client](#javascript-api-client)
- [Sample Frontend Pages](#sample-frontend-pages)
- [Error Handling](#error-handling)
- [Best Practices](#best-practices)

## API Overview

**Base URL**: `http://localhost:3644`
**API Version**: `/api/v1`
**Content-Type**: `application/json`

### Available Endpoints

| Entity | Endpoint | Methods | Description |
|--------|----------|---------|-------------|
| Restaurants | `/api/v1/restaurants` | GET, POST | List/Create restaurants |
| | `/api/v1/restaurants/:id` | GET, PUT, DELETE | Get/Update/Delete restaurant |
| Menu Items | `/api/v1/menu-items` | GET, POST | List/Create menu items |
| | `/api/v1/menu-items/:id` | GET, PUT, DELETE | Get/Update/Delete menu item |
| Customers | `/api/v1/customers` | GET, POST | List/Create customers |
| | `/api/v1/customers/:id` | GET, PUT, DELETE | Get/Update/Delete customer |
| Orders | `/api/v1/orders` | GET, POST | List/Create orders |
| | `/api/v1/orders/:id` | GET, DELETE | Get/Delete order |
| | `/api/v1/orders/:id/status` | PATCH | Update order status |

## Authentication & CORS

The API currently **does not require authentication** and has **CORS enabled** for development. You can make requests directly from the browser.

## Basic HTML Setup

### 1. Basic HTML Template

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Restaurant Management</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .card { border: 1px solid #ddd; padding: 20px; margin: 10px 0; border-radius: 5px; }
        .form-group { margin: 10px 0; }
        label { display: block; margin-bottom: 5px; font-weight: bold; }
        input, select, textarea { width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 3px; }
        button { background: #007bff; color: white; padding: 10px 20px; border: none; border-radius: 3px; cursor: pointer; }
        button:hover { background: #0056b3; }
        .error { color: red; margin: 10px 0; }
        .success { color: green; margin: 10px 0; }
        table { width: 100%; border-collapse: collapse; margin: 20px 0; }
        th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
        th { background-color: #f2f2f2; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Restaurant Management System</h1>
        <!-- Your content here -->
    </div>
    <script src="api-client.js"></script>
    <script src="app.js"></script>
</body>
</html>
```

## JavaScript API Client

### 2. Create `api-client.js`

```javascript
class RestaurantAPI {
    constructor(baseURL = 'http://localhost:3644') {
        this.baseURL = baseURL;
    }

    // Generic request method
    async request(endpoint, options = {}) {
        const url = `${this.baseURL}${endpoint}`;
        const config = {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            },
            ...options
        };

        try {
            const response = await fetch(url, config);
            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.error || `HTTP ${response.status}`);
            }
            
            return data;
        } catch (error) {
            console.error('API Request failed:', error);
            throw error;
        }
    }

    // Restaurant methods
    async getRestaurants() {
        return this.request('/api/v1/restaurants');
    }

    async getRestaurant(id) {
        return this.request(`/api/v1/restaurants/${id}`);
    }

    async createRestaurant(restaurant) {
        return this.request('/api/v1/restaurants', {
            method: 'POST',
            body: JSON.stringify(restaurant)
        });
    }

    async updateRestaurant(id, restaurant) {
        return this.request(`/api/v1/restaurants/${id}`, {
            method: 'PUT',
            body: JSON.stringify(restaurant)
        });
    }

    async deleteRestaurant(id) {
        return this.request(`/api/v1/restaurants/${id}`, {
            method: 'DELETE'
        });
    }

    // Menu Item methods
    async getMenuItems(restaurantId = null) {
        const params = restaurantId ? `?restaurant_id=${restaurantId}` : '';
        return this.request(`/api/v1/menu-items${params}`);
    }

    async getMenuItem(id) {
        return this.request(`/api/v1/menu-items/${id}`);
    }

    async createMenuItem(menuItem) {
        return this.request('/api/v1/menu-items', {
            method: 'POST',
            body: JSON.stringify(menuItem)
        });
    }

    async updateMenuItem(id, menuItem) {
        return this.request(`/api/v1/menu-items/${id}`, {
            method: 'PUT',
            body: JSON.stringify(menuItem)
        });
    }

    async deleteMenuItem(id) {
        return this.request(`/api/v1/menu-items/${id}`, {
            method: 'DELETE'
        });
    }

    // Customer methods
    async getCustomers() {
        return this.request('/api/v1/customers');
    }

    async getCustomer(id) {
        return this.request(`/api/v1/customers/${id}`);
    }

    async createCustomer(customer) {
        return this.request('/api/v1/customers', {
            method: 'POST',
            body: JSON.stringify(customer)
        });
    }

    async updateCustomer(id, customer) {
        return this.request(`/api/v1/customers/${id}`, {
            method: 'PUT',
            body: JSON.stringify(customer)
        });
    }

    async deleteCustomer(id) {
        return this.request(`/api/v1/customers/${id}`, {
            method: 'DELETE'
        });
    }

    // Order methods
    async getOrders(customerId = null, restaurantId = null) {
        const params = new URLSearchParams();
        if (customerId) params.append('customer_id', customerId);
        if (restaurantId) params.append('restaurant_id', restaurantId);
        const queryString = params.toString() ? `?${params.toString()}` : '';
        return this.request(`/api/v1/orders${queryString}`);
    }

    async getOrder(id) {
        return this.request(`/api/v1/orders/${id}`);
    }

    async createOrder(order) {
        return this.request('/api/v1/orders', {
            method: 'POST',
            body: JSON.stringify(order)
        });
    }

    async updateOrderStatus(id, status) {
        return this.request(`/api/v1/orders/${id}/status`, {
            method: 'PATCH',
            body: JSON.stringify({ status })
        });
    }

    async deleteOrder(id) {
        return this.request(`/api/v1/orders/${id}`, {
            method: 'DELETE'
        });
    }

    // Health check
    async healthCheck() {
        return this.request('/health');
    }
}

// Create global API instance
const api = new RestaurantAPI();
```

## Sample Frontend Pages

### 3. Restaurant Management Page

```html
<!-- restaurants.html -->
<div class="card">
    <h2>Restaurants</h2>
    
    <!-- Add Restaurant Form -->
    <div class="card">
        <h3>Add New Restaurant</h3>
        <form id="restaurantForm">
            <div class="form-group">
                <label for="name">Name:</label>
                <input type="text" id="name" name="name" required>
            </div>
            <div class="form-group">
                <label for="address">Address:</label>
                <input type="text" id="address" name="address">
            </div>
            <div class="form-group">
                <label for="phone">Phone:</label>
                <input type="tel" id="phone" name="phone">
            </div>
            <div class="form-group">
                <label for="email">Email:</label>
                <input type="email" id="email" name="email">
            </div>
            <div class="form-group">
                <label for="cuisine_type">Cuisine Type:</label>
                <input type="text" id="cuisine_type" name="cuisine_type">
            </div>
            <button type="submit">Add Restaurant</button>
        </form>
    </div>

    <!-- Restaurants List -->
    <div id="restaurantsList"></div>
</div>

<script>
// Restaurant management JavaScript
document.addEventListener('DOMContentLoaded', function() {
    loadRestaurants();
    
    // Handle form submission
    document.getElementById('restaurantForm').addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const formData = new FormData(e.target);
        const restaurant = Object.fromEntries(formData.entries());
        
        try {
            await api.createRestaurant(restaurant);
            showMessage('Restaurant added successfully!', 'success');
            e.target.reset();
            loadRestaurants();
        } catch (error) {
            showMessage('Error: ' + error.message, 'error');
        }
    });
});

async function loadRestaurants() {
    try {
        const restaurants = await api.getRestaurants();
        displayRestaurants(restaurants);
    } catch (error) {
        showMessage('Error loading restaurants: ' + error.message, 'error');
    }
}

function displayRestaurants(restaurants) {
    const container = document.getElementById('restaurantsList');
    
    if (restaurants.length === 0) {
        container.innerHTML = '<p>No restaurants found.</p>';
        return;
    }
    
    const html = `
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Address</th>
                    <th>Phone</th>
                    <th>Email</th>
                    <th>Cuisine Type</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                ${restaurants.map(restaurant => `
                    <tr>
                        <td>${restaurant.id}</td>
                        <td>${restaurant.name}</td>
                        <td>${restaurant.address || 'N/A'}</td>
                        <td>${restaurant.phone || 'N/A'}</td>
                        <td>${restaurant.email || 'N/A'}</td>
                        <td>${restaurant.cuisine_type || 'N/A'}</td>
                        <td>
                            <button onclick="editRestaurant(${restaurant.id})">Edit</button>
                            <button onclick="deleteRestaurant(${restaurant.id})">Delete</button>
                        </td>
                    </tr>
                `).join('')}
            </tbody>
        </table>
    `;
    
    container.innerHTML = html;
}

async function deleteRestaurant(id) {
    if (confirm('Are you sure you want to delete this restaurant?')) {
        try {
            await api.deleteRestaurant(id);
            showMessage('Restaurant deleted successfully!', 'success');
            loadRestaurants();
        } catch (error) {
            showMessage('Error: ' + error.message, 'error');
        }
    }
}

function showMessage(message, type) {
    const messageDiv = document.createElement('div');
    messageDiv.className = type;
    messageDiv.textContent = message;
    
    const container = document.querySelector('.container');
    container.insertBefore(messageDiv, container.firstChild);
    
    setTimeout(() => messageDiv.remove(), 5000);
}
</script>
```

### 4. Order Management Page

```html
<!-- orders.html -->
<div class="card">
    <h2>Order Management</h2>
    
    <!-- Create Order Form -->
    <div class="card">
        <h3>Create New Order</h3>
        <form id="orderForm">
            <div class="form-group">
                <label for="customer_id">Customer:</label>
                <select id="customer_id" name="customer_id" required></select>
            </div>
            <div class="form-group">
                <label for="restaurant_id">Restaurant:</label>
                <select id="restaurant_id" name="restaurant_id" required></select>
            </div>
            <div class="form-group">
                <label for="delivery_address">Delivery Address:</label>
                <input type="text" id="delivery_address" name="delivery_address">
            </div>
            <div class="form-group">
                <label for="notes">Notes:</label>
                <textarea id="notes" name="notes" rows="3"></textarea>
            </div>
            
            <h4>Order Items</h4>
            <div id="orderItems">
                <div class="form-group">
                    <label for="menu_item_id_1">Menu Item:</label>
                    <select id="menu_item_id_1" name="menu_item_id_1" required></select>
                    <label for="quantity_1">Quantity:</label>
                    <input type="number" id="quantity_1" name="quantity_1" min="1" value="1" required>
                </div>
            </div>
            
            <button type="button" onclick="addOrderItem()">Add Item</button>
            <button type="submit">Create Order</button>
        </form>
    </div>

    <!-- Orders List -->
    <div id="ordersList"></div>
</div>

<script>
let orderItemCount = 1;

document.addEventListener('DOMContentLoaded', function() {
    loadOrdersPage();
    
    document.getElementById('orderForm').addEventListener('submit', async function(e) {
        e.preventDefault();
        await createOrder();
    });
    
    // Load restaurant change handler
    document.getElementById('restaurant_id').addEventListener('change', function() {
        loadMenuItems(this.value);
    });
});

async function loadOrdersPage() {
    await Promise.all([
        loadCustomers(),
        loadRestaurants(),
        loadOrders()
    ]);
}

async function loadCustomers() {
    try {
        const customers = await api.getCustomers();
        const select = document.getElementById('customer_id');
        select.innerHTML = '<option value="">Select a customer</option>' +
            customers.map(c => `<option value="${c.id}">${c.name}</option>`).join('');
    } catch (error) {
        console.error('Error loading customers:', error);
    }
}

async function loadRestaurants() {
    try {
        const restaurants = await api.getRestaurants();
        const select = document.getElementById('restaurant_id');
        select.innerHTML = '<option value="">Select a restaurant</option>' +
            restaurants.map(r => `<option value="${r.id}">${r.name}</option>`).join('');
    } catch (error) {
        console.error('Error loading restaurants:', error);
    }
}

async function loadMenuItems(restaurantId) {
    if (!restaurantId) return;
    
    try {
        const menuItems = await api.getMenuItems(restaurantId);
        const selects = document.querySelectorAll('[name^="menu_item_id_"]');
        
        selects.forEach(select => {
            select.innerHTML = '<option value="">Select a menu item</option>' +
                menuItems.map(item => 
                    `<option value="${item.id}">${item.name} - $${item.price}</option>`
                ).join('');
        });
    } catch (error) {
        console.error('Error loading menu items:', error);
    }
}

function addOrderItem() {
    orderItemCount++;
    const container = document.getElementById('orderItems');
    
    const div = document.createElement('div');
    div.className = 'form-group';
    div.innerHTML = `
        <label for="menu_item_id_${orderItemCount}">Menu Item:</label>
        <select id="menu_item_id_${orderItemCount}" name="menu_item_id_${orderItemCount}" required>
            <option value="">Select a menu item</option>
        </select>
        <label for="quantity_${orderItemCount}">Quantity:</label>
        <input type="number" id="quantity_${orderItemCount}" name="quantity_${orderItemCount}" min="1" value="1" required>
        <button type="button" onclick="removeOrderItem(this)">Remove</button>
    `;
    
    container.appendChild(div);
    
    // Load menu items for the new select
    const restaurantId = document.getElementById('restaurant_id').value;
    if (restaurantId) {
        loadMenuItems(restaurantId);
    }
}

function removeOrderItem(button) {
    button.parentElement.remove();
}

async function createOrder() {
    const formData = new FormData(document.getElementById('orderForm'));
    
    // Extract basic order data
    const order = {
        customer_id: parseInt(formData.get('customer_id')),
        restaurant_id: parseInt(formData.get('restaurant_id')),
        delivery_address: formData.get('delivery_address'),
        notes: formData.get('notes'),
        items: []
    };
    
    // Extract order items
    for (let i = 1; i <= orderItemCount; i++) {
        const menuItemId = formData.get(`menu_item_id_${i}`);
        const quantity = formData.get(`quantity_${i}`);
        
        if (menuItemId && quantity) {
            order.items.push({
                menu_item_id: parseInt(menuItemId),
                quantity: parseInt(quantity)
            });
        }
    }
    
    if (order.items.length === 0) {
        showMessage('Please add at least one item to the order', 'error');
        return;
    }
    
    try {
        await api.createOrder(order);
        showMessage('Order created successfully!', 'success');
        document.getElementById('orderForm').reset();
        orderItemCount = 1;
        document.getElementById('orderItems').innerHTML = `
            <div class="form-group">
                <label for="menu_item_id_1">Menu Item:</label>
                <select id="menu_item_id_1" name="menu_item_id_1" required>
                    <option value="">Select a menu item</option>
                </select>
                <label for="quantity_1">Quantity:</label>
                <input type="number" id="quantity_1" name="quantity_1" min="1" value="1" required>
            </div>
        `;
        loadOrders();
    } catch (error) {
        showMessage('Error creating order: ' + error.message, 'error');
    }
}

async function loadOrders() {
    try {
        const orders = await api.getOrders();
        displayOrders(orders);
    } catch (error) {
        showMessage('Error loading orders: ' + error.message, 'error');
    }
}

function displayOrders(orders) {
    const container = document.getElementById('ordersList');
    
    if (orders.length === 0) {
        container.innerHTML = '<p>No orders found.</p>';
        return;
    }
    
    const html = `
        <h3>Orders</h3>
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Customer ID</th>
                    <th>Restaurant ID</th>
                    <th>Total</th>
                    <th>Status</th>
                    <th>Order Date</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                ${orders.map(order => `
                    <tr>
                        <td>${order.id}</td>
                        <td>${order.customer_id}</td>
                        <td>${order.restaurant_id}</td>
                        <td>$${order.total_amount.toFixed(2)}</td>
                        <td>
                            <select onchange="updateOrderStatus(${order.id}, this.value)">
                                <option value="pending" ${order.status === 'pending' ? 'selected' : ''}>Pending</option>
                                <option value="confirmed" ${order.status === 'confirmed' ? 'selected' : ''}>Confirmed</option>
                                <option value="preparing" ${order.status === 'preparing' ? 'selected' : ''}>Preparing</option>
                                <option value="ready" ${order.status === 'ready' ? 'selected' : ''}>Ready</option>
                                <option value="delivered" ${order.status === 'delivered' ? 'selected' : ''}>Delivered</option>
                                <option value="cancelled" ${order.status === 'cancelled' ? 'selected' : ''}>Cancelled</option>
                            </select>
                        </td>
                        <td>${new Date(order.order_date).toLocaleDateString()}</td>
                        <td>
                            <button onclick="viewOrder(${order.id})">View</button>
                            <button onclick="deleteOrder(${order.id})">Delete</button>
                        </td>
                    </tr>
                `).join('')}
            </tbody>
        </table>
    `;
    
    container.innerHTML = html;
}

async function updateOrderStatus(orderId, newStatus) {
    try {
        await api.updateOrderStatus(orderId, newStatus);
        showMessage('Order status updated successfully!', 'success');
    } catch (error) {
        showMessage('Error updating order status: ' + error.message, 'error');
        loadOrders(); // Reload to reset the select
    }
}

async function deleteOrder(id) {
    if (confirm('Are you sure you want to delete this order?')) {
        try {
            await api.deleteOrder(id);
            showMessage('Order deleted successfully!', 'success');
            loadOrders();
        } catch (error) {
            showMessage('Error deleting order: ' + error.message, 'error');
        }
    }
}

async function viewOrder(id) {
    try {
        const order = await api.getOrder(id);
        
        let itemsHtml = '<h4>Order Items:</h4><ul>';
        order.items.forEach(item => {
            itemsHtml += `<li>${item.menu_item.name} - Quantity: ${item.quantity} - $${item.unit_price.toFixed(2)}</li>`;
        });
        itemsHtml += '</ul>';
        
        const orderDetails = `
            <div class="card">
                <h3>Order #${order.id}</h3>
                <p><strong>Customer ID:</strong> ${order.customer_id}</p>
                <p><strong>Restaurant ID:</strong> ${order.restaurant_id}</p>
                <p><strong>Status:</strong> ${order.status}</p>
                <p><strong>Total:</strong> $${order.total_amount.toFixed(2)}</p>
                <p><strong>Order Date:</strong> ${new Date(order.order_date).toLocaleString()}</p>
                <p><strong>Delivery Address:</strong> ${order.delivery_address || 'N/A'}</p>
                <p><strong>Notes:</strong> ${order.notes || 'None'}</p>
                ${itemsHtml}
                <button onclick="this.parentElement.remove()">Close</button>
            </div>
        `;
        
        const container = document.querySelector('.container');
        container.insertAdjacentHTML('beforeend', orderDetails);
    } catch (error) {
        showMessage('Error loading order details: ' + error.message, 'error');
    }
}

function showMessage(message, type) {
    const messageDiv = document.createElement('div');
    messageDiv.className = type;
    messageDiv.textContent = message;
    
    const container = document.querySelector('.container');
    container.insertBefore(messageDiv, container.firstChild);
    
    setTimeout(() => messageDiv.remove(), 5000);
}
</script>
```

## Error Handling

### 5. Comprehensive Error Handling

```javascript
// Enhanced API client with better error handling
class RestaurantAPI {
    // ... (previous code)
    
    async request(endpoint, options = {}) {
        const url = `${this.baseURL}${endpoint}`;
        const config = {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            },
            ...options
        };

        try {
            const response = await fetch(url, config);
            
            // Handle different response types
            let data;
            const contentType = response.headers.get('content-type');
            
            if (contentType && contentType.includes('application/json')) {
                data = await response.json();
            } else {
                data = await response.text();
            }
            
            if (!response.ok) {
                // Handle different error types
                const errorMessage = typeof data === 'object' && data.error 
                    ? data.error 
                    : `HTTP ${response.status}: ${response.statusText}`;
                
                throw new APIError(errorMessage, response.status, data);
            }
            
            return data;
        } catch (error) {
            if (error instanceof APIError) {
                throw error;
            }
            
            // Handle network errors
            if (error.name === 'TypeError' && error.message.includes('fetch')) {
                throw new APIError('Network error: Unable to connect to server', 0);
            }
            
            console.error('API Request failed:', error);
            throw new APIError(error.message || 'Unknown error occurred');
        }
    }
}

// Custom error class
class APIError extends Error {
    constructor(message, status = 0, data = null) {
        super(message);
        this.name = 'APIError';
        this.status = status;
        this.data = data;
    }
}

// Global error handler
window.addEventListener('unhandledrejection', function(event) {
    console.error('Unhandled promise rejection:', event.reason);
    showMessage('An unexpected error occurred. Please try again.', 'error');
});
```

## Best Practices

### 6. Frontend Best Practices

1. **Input Validation**
```javascript
function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(email);
}

function validatePhone(phone) {
    const re = /^\+?[\d\s\-\(\)]+$/;
    return re.test(phone);
}

function validateForm(formData) {
    const errors = [];
    
    if (!formData.get('name') || formData.get('name').trim().length < 2) {
        errors.push('Name must be at least 2 characters long');
    }
    
    const email = formData.get('email');
    if (email && !validateEmail(email)) {
        errors.push('Please enter a valid email address');
    }
    
    return errors;
}
```

2. **Loading States**
```javascript
function showLoading(element) {
    element.innerHTML = '<div>Loading...</div>';
    element.style.opacity = '0.6';
}

function hideLoading(element) {
    element.style.opacity = '1';
}

// Usage
async function loadRestaurants() {
    const container = document.getElementById('restaurantsList');
    showLoading(container);
    
    try {
        const restaurants = await api.getRestaurants();
        displayRestaurants(restaurants);
    } catch (error) {
        showMessage('Error loading restaurants: ' + error.message, 'error');
    } finally {
        hideLoading(container);
    }
}
```

3. **Debounced Search**
```javascript
function debounce(func, wait) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

// Search functionality
const searchRestaurants = debounce(async function(query) {
    if (query.length < 2) return;
    
    try {
        const restaurants = await api.getRestaurants();
        const filtered = restaurants.filter(r => 
            r.name.toLowerCase().includes(query.toLowerCase()) ||
            (r.cuisine_type && r.cuisine_type.toLowerCase().includes(query.toLowerCase()))
        );
        displayRestaurants(filtered);
    } catch (error) {
        showMessage('Error searching restaurants: ' + error.message, 'error');
    }
}, 300);
```

4. **Local Storage Cache**
```javascript
class CachedAPI extends RestaurantAPI {
    constructor(baseURL) {
        super(baseURL);
        this.cache = new Map();
        this.cacheTimeout = 5 * 60 * 1000; // 5 minutes
    }
    
    async getRestaurants() {
        const cacheKey = 'restaurants';
        const cached = this.getFromCache(cacheKey);
        
        if (cached) {
            return cached;
        }
        
        const data = await super.getRestaurants();
        this.setCache(cacheKey, data);
        return data;
    }
    
    getFromCache(key) {
        const cached = this.cache.get(key);
        if (cached && Date.now() - cached.timestamp < this.cacheTimeout) {
            return cached.data;
        }
        return null;
    }
    
    setCache(key, data) {
        this.cache.set(key, {
            data,
            timestamp: Date.now()
        });
    }
    
    clearCache() {
        this.cache.clear();
    }
}
```

### Quick Start Files

Create these files in your frontend project:

1. `index.html` - Main page with navigation
2. `api-client.js` - API wrapper class
3. `restaurants.html` - Restaurant management
4. `menu-items.html` - Menu item management  
5. `customers.html` - Customer management
6. `orders.html` - Order management
7. `styles.css` - Custom styles
8. `app.js` - Main application logic

### Testing Your Frontend

1. Start the Go backend: `go run .`
2. Open your HTML files in a web browser
3. Check browser console for any errors
4. Test all CRUD operations
5. Verify error handling with invalid data

This guide provides everything needed to create a functional frontend for the Restaurant API! 