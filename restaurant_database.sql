-- Restaurant Management Database Setup Script
-- Run this script in your MySQL database to create all tables and demo data

-- Create database (uncomment if needed)
-- CREATE DATABASE IF NOT EXISTS restuarant CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- USE restuarant;

-- Drop existing tables in correct order (to handle foreign key constraints)
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS restaurants;
SET FOREIGN_KEY_CHECKS = 1;

-- Create restaurants table
CREATE TABLE restaurants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(500),
    phone VARCHAR(20),
    email VARCHAR(255),
    cuisine_type VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create customers table
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(20),
    address VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create menu_items table
CREATE TABLE menu_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    restaurant_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category VARCHAR(100),
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE
);

-- Create orders table
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    restaurant_id INT NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status ENUM('pending', 'confirmed', 'preparing', 'ready', 'delivered', 'cancelled') DEFAULT 'pending',
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    delivery_address VARCHAR(500),
    notes TEXT,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE RESTRICT,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE RESTRICT
);

-- Create order_items table
CREATE TABLE order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT NOT NULL,
    menu_item_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (menu_item_id) REFERENCES menu_items(id) ON DELETE RESTRICT
);

-- Insert demo restaurants
INSERT INTO restaurants (name, address, phone, email, cuisine_type) VALUES
('Pizza Palace', '123 Main St, Downtown', '+1-555-0101', 'info@pizzapalace.com', 'Italian'),
('Burger Barn', '456 Oak Ave, Midtown', '+1-555-0102', 'orders@burgerbarn.com', 'American'),
('Sushi Zen', '789 Pine Rd, Uptown', '+1-555-0103', 'hello@sushizen.com', 'Japanese'),
('Taco Fiesta', '321 Elm St, Downtown', '+1-555-0104', 'contact@tacofiesta.com', 'Mexican'),
('Thai Garden', '654 Maple Ave, Eastside', '+1-555-0105', 'info@thaigarden.com', 'Thai');

-- Insert demo customers
INSERT INTO customers (name, email, phone, address) VALUES
('John Doe', 'john.doe@example.com', '+1-555-1001', '100 Customer St, Residential Area'),
('Jane Smith', 'jane.smith@example.com', '+1-555-1002', '200 Buyer Ave, Suburb'),
('Bob Wilson', 'bob.wilson@example.com', '+1-555-1003', '300 Client Rd, Downtown'),
('Alice Johnson', 'alice.johnson@example.com', '+1-555-1004', '400 User Blvd, Midtown'),
('Charlie Brown', 'charlie.brown@example.com', '+1-555-1005', '500 Guest Lane, Uptown');

-- Insert demo menu items for Pizza Palace (restaurant_id = 1)
INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available) VALUES
(1, 'Margherita Pizza', 'Fresh tomato sauce, mozzarella cheese, fresh basil', 12.99, 'Pizza', TRUE),
(1, 'Pepperoni Pizza', 'Tomato sauce, mozzarella, pepperoni slices', 14.99, 'Pizza', TRUE),
(1, 'Supreme Pizza', 'Pepperoni, sausage, bell peppers, onions, mushrooms', 17.99, 'Pizza', TRUE),
(1, 'Caesar Salad', 'Romaine lettuce, croutons, parmesan, caesar dressing', 8.99, 'Salad', TRUE),
(1, 'Garlic Bread', 'Fresh baked bread with garlic butter', 4.99, 'Appetizer', TRUE),
(1, 'Tiramisu', 'Classic Italian dessert with coffee and mascarpone', 6.99, 'Dessert', TRUE);

-- Insert demo menu items for Burger Barn (restaurant_id = 2)
INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available) VALUES
(2, 'Classic Burger', 'Beef patty, lettuce, tomato, onion, pickles', 9.99, 'Burger', TRUE),
(2, 'Cheeseburger', 'Classic burger with american cheese', 10.99, 'Burger', TRUE),
(2, 'BBQ Bacon Burger', 'Beef patty, bacon, BBQ sauce, onion rings', 12.99, 'Burger', TRUE),
(2, 'Chicken Sandwich', 'Grilled chicken breast, mayo, lettuce', 8.99, 'Sandwich', TRUE),
(2, 'French Fries', 'Crispy golden fries', 3.99, 'Side', TRUE),
(2, 'Onion Rings', 'Beer-battered onion rings', 4.99, 'Side', TRUE),
(2, 'Milkshake', 'Vanilla, chocolate, or strawberry', 4.99, 'Beverage', TRUE);

-- Insert demo menu items for Sushi Zen (restaurant_id = 3)
INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available) VALUES
(3, 'Salmon Roll', 'Fresh salmon, avocado, cucumber', 8.99, 'Roll', TRUE),
(3, 'Tuna Roll', 'Fresh tuna, avocado', 9.99, 'Roll', TRUE),
(3, 'California Roll', 'Crab, avocado, cucumber', 7.99, 'Roll', TRUE),
(3, 'Salmon Sashimi', 'Fresh salmon slices (6 pieces)', 12.99, 'Sashimi', TRUE),
(3, 'Tuna Sashimi', 'Fresh tuna slices (6 pieces)', 14.99, 'Sashimi', TRUE),
(3, 'Miso Soup', 'Traditional soybean soup', 3.99, 'Soup', TRUE),
(3, 'Edamame', 'Steamed and salted soybeans', 4.99, 'Appetizer', TRUE);

-- Insert demo menu items for Taco Fiesta (restaurant_id = 4)
INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available) VALUES
(4, 'Beef Tacos', 'Seasoned ground beef, lettuce, cheese (3 tacos)', 8.99, 'Tacos', TRUE),
(4, 'Chicken Tacos', 'Grilled chicken, salsa, cheese (3 tacos)', 9.99, 'Tacos', TRUE),
(4, 'Fish Tacos', 'Grilled fish, cabbage slaw, lime (3 tacos)', 11.99, 'Tacos', TRUE),
(4, 'Beef Burrito', 'Large flour tortilla with beef, beans, rice', 10.99, 'Burrito', TRUE),
(4, 'Chicken Quesadilla', 'Grilled chicken and cheese in tortilla', 8.99, 'Quesadilla', TRUE),
(4, 'Guacamole & Chips', 'Fresh guacamole with tortilla chips', 5.99, 'Appetizer', TRUE),
(4, 'Churros', 'Fried pastry with cinnamon sugar', 4.99, 'Dessert', TRUE);

-- Insert demo menu items for Thai Garden (restaurant_id = 5)
INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available) VALUES
(5, 'Pad Thai', 'Stir-fried rice noodles with shrimp or chicken', 12.99, 'Noodles', TRUE),
(5, 'Green Curry', 'Coconut curry with vegetables and choice of protein', 13.99, 'Curry', TRUE),
(5, 'Tom Yum Soup', 'Spicy and sour soup with shrimp', 8.99, 'Soup', TRUE),
(5, 'Spring Rolls', 'Fresh vegetables wrapped in rice paper (4 rolls)', 6.99, 'Appetizer', TRUE),
(5, 'Mango Sticky Rice', 'Sweet sticky rice with fresh mango', 5.99, 'Dessert', TRUE),
(5, 'Thai Iced Tea', 'Sweet tea with condensed milk', 3.99, 'Beverage', TRUE);

-- Insert demo orders
INSERT INTO orders (customer_id, restaurant_id, total_amount, status, delivery_address, notes) VALUES
(1, 1, 21.98, 'confirmed', '100 Customer St, Residential Area', 'Please ring doorbell, leave at door'),
(2, 2, 19.97, 'preparing', '200 Buyer Ave, Suburb', 'Extra sauce on burger, no pickles'),
(3, 3, 25.97, 'ready', '300 Client Rd, Downtown', 'Hold the wasabi, extra ginger'),
(4, 4, 18.98, 'delivered', '400 User Blvd, Midtown', 'Delivered successfully'),
(5, 5, 16.98, 'pending', '500 Guest Lane, Uptown', 'Call when arriving'),
(1, 2, 15.98, 'cancelled', '100 Customer St, Residential Area', 'Customer cancelled order');

-- Insert demo order items
INSERT INTO order_items (order_id, menu_item_id, quantity, unit_price) VALUES
-- Order 1 (Pizza Palace): Margherita Pizza + Caesar Salad
(1, 1, 1, 12.99),
(1, 4, 1, 8.99),

-- Order 2 (Burger Barn): Cheeseburger + French Fries + Milkshake
(2, 8, 1, 10.99),
(2, 11, 1, 3.99),
(2, 13, 1, 4.99),

-- Order 3 (Sushi Zen): Salmon Roll + Tuna Sashimi + Miso Soup
(3, 14, 1, 8.99),
(3, 18, 1, 14.99),
(3, 19, 1, 3.99),

-- Order 4 (Taco Fiesta): Chicken Tacos + Guacamole & Chips
(4, 22, 1, 9.99),
(4, 26, 1, 5.99),

-- Order 5 (Thai Garden): Pad Thai + Thai Iced Tea
(5, 28, 1, 12.99),
(5, 33, 1, 3.99),

-- Order 6 (Burger Barn - Cancelled): Classic Burger + French Fries
(6, 7, 1, 9.99),
(6, 11, 1, 3.99);

-- Create indexes for better performance
CREATE INDEX idx_menu_items_restaurant_id ON menu_items(restaurant_id);
CREATE INDEX idx_menu_items_category ON menu_items(category);
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_orders_restaurant_id ON orders(restaurant_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_date ON orders(order_date);
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_menu_item_id ON order_items(menu_item_id);
CREATE INDEX idx_customers_email ON customers(email);

-- Display summary of created data
SELECT 'Database Setup Complete!' as Status;

SELECT 
    'Restaurants' as Entity,
    COUNT(*) as Count
FROM restaurants
UNION ALL
SELECT 
    'Customers' as Entity,
    COUNT(*) as Count
FROM customers
UNION ALL
SELECT 
    'Menu Items' as Entity,
    COUNT(*) as Count
FROM menu_items
UNION ALL
SELECT 
    'Orders' as Entity,
    COUNT(*) as Count
FROM orders
UNION ALL
SELECT 
    'Order Items' as Entity,
    COUNT(*) as Count
FROM order_items;

-- Show restaurants with menu item counts
SELECT 
    r.id,
    r.name as restaurant_name,
    r.cuisine_type,
    COUNT(m.id) as menu_items_count
FROM restaurants r
LEFT JOIN menu_items m ON r.id = m.restaurant_id
GROUP BY r.id, r.name, r.cuisine_type
ORDER BY r.id;

-- Show orders summary
SELECT 
    o.id as order_id,
    c.name as customer_name,
    r.name as restaurant_name,
    o.total_amount,
    o.status,
    COUNT(oi.id) as items_count
FROM orders o
JOIN customers c ON o.customer_id = c.id
JOIN restaurants r ON o.restaurant_id = r.id
LEFT JOIN order_items oi ON o.id = oi.order_id
GROUP BY o.id, c.name, r.name, o.total_amount, o.status
ORDER BY o.id; 