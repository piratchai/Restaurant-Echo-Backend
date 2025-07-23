# Restaurant Backend API

A complete RESTful API for restaurant management with CRUD operations built using Go Echo framework and MySQL.

## Features

- **Restaurant Management**: Create, read, update, delete restaurants
- **Menu Management**: Manage menu items for each restaurant
- **Customer Management**: Handle customer data and profiles
- **Order Management**: Process orders with multiple items and status tracking
- **Database Integration**: MySQL database with proper relationships
- **RESTful API**: Clean API design with JSON responses
- **Error Handling**: Comprehensive error handling and validation

## Database Schema

### Tables

1. **restaurants**: Restaurant information
2. **menu_items**: Menu items belonging to restaurants
3. **customers**: Customer profiles
4. **orders**: Order headers with customer and restaurant info
5. **order_items**: Individual items in each order

## API Endpoints

### Root & Health
- `GET /` - API information
- `GET /health` - Health check

### Restaurants
- `POST /api/v1/restaurants` - Create restaurant
- `GET /api/v1/restaurants` - List all restaurants
- `GET /api/v1/restaurants/:id` - Get restaurant by ID
- `PUT /api/v1/restaurants/:id` - Update restaurant
- `DELETE /api/v1/restaurants/:id` - Delete restaurant

### Menu Items
- `POST /api/v1/menu-items` - Create menu item
- `GET /api/v1/menu-items` - List menu items (supports ?restaurant_id=X filter)
- `GET /api/v1/menu-items/:id` - Get menu item by ID
- `PUT /api/v1/menu-items/:id` - Update menu item
- `DELETE /api/v1/menu-items/:id` - Delete menu item

### Customers
- `POST /api/v1/customers` - Create customer
- `GET /api/v1/customers` - List all customers
- `GET /api/v1/customers/:id` - Get customer by ID
- `PUT /api/v1/customers/:id` - Update customer
- `DELETE /api/v1/customers/:id` - Delete customer

### Orders
- `POST /api/v1/orders` - Create order with items
- `GET /api/v1/orders` - List orders (supports ?customer_id=X and ?restaurant_id=X filters)
- `GET /api/v1/orders/:id` - Get order by ID with items
- `PATCH /api/v1/orders/:id/status` - Update order status
- `DELETE /api/v1/orders/:id` - Delete order

## Sample Requests

### Create Restaurant
```bash
curl -X POST http://localhost:3644/api/v1/restaurants \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pizza Palace",
    "address": "123 Main St",
    "phone": "+1-555-0101",
    "email": "info@pizzapalace.com",
    "cuisine_type": "Italian"
  }'
```

### Create Menu Item
```bash
curl -X POST http://localhost:3644/api/v1/menu-items \
  -H "Content-Type: application/json" \
  -d '{
    "restaurant_id": 1,
    "name": "Margherita Pizza",
    "description": "Fresh tomato sauce, mozzarella, basil",
    "price": 12.99,
    "category": "Pizza",
    "is_available": true
  }'
```

### Create Customer
```bash
curl -X POST http://localhost:3644/api/v1/customers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1-555-1001",
    "address": "100 Customer St"
  }'
```

### Create Order
```bash
curl -X POST http://localhost:3644/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "restaurant_id": 1,
    "delivery_address": "100 Customer St",
    "notes": "Ring doorbell",
    "items": [
      {"menu_item_id": 1, "quantity": 2},
      {"menu_item_id": 2, "quantity": 1}
    ]
  }'
```

### Update Order Status
```bash
curl -X PATCH http://localhost:3644/api/v1/orders/1/status \
  -H "Content-Type: application/json" \
  -d '{"status": "confirmed"}'
```

## Order Status Values
- `pending` - Order placed, awaiting confirmation
- `confirmed` - Order confirmed by restaurant
- `preparing` - Order being prepared
- `ready` - Order ready for pickup/delivery
- `delivered` - Order delivered
- `cancelled` - Order cancelled

## Setup & Installation

1. **Install Dependencies**
   ```bash
   go mod tidy
   ```

2. **Configure Database**
   - Update the database connection string in `database.go`
   - Ensure MySQL is running and accessible

3. **Run the Server**
   ```bash
   go run .
   ```

4. **Test the API**
   ```bash
   curl http://localhost:3644/health
   ```

## Project Structure

```
echo_demo1/
├── go.mod              # Go module file
├── go.sum              # Go dependencies
├── server.go           # Main server with routes
├── models.go           # Data models and structs
├── database.go         # Database connection
├── handlers.go         # Restaurant & menu item handlers
├── customer_handlers.go # Customer CRUD handlers
├── order_handlers.go   # Order CRUD handlers
└── README.md          # This file
```

## Database Design

The database uses foreign key relationships:
- `menu_items.restaurant_id` → `restaurants.id`
- `orders.customer_id` → `customers.id`
- `orders.restaurant_id` → `restaurants.id`
- `order_items.order_id` → `orders.id`
- `order_items.menu_item_id` → `menu_items.id`

Orders automatically calculate total amounts and support cascading deletes for order items. 