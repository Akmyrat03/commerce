CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    image VARCHAR(255),
    price DECIMAL(10, 2) NOT NULL,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE
);

-- CREATE TABLE IF NOT EXISTS orders (
--     id SERIAL PRIMARY KEY,
--     user_id INT REFERENCES users(id),
--     total_amount DECIMAL(10, 2) NOT NULL,
--     status VARCHAR(50),
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- CREATE TABLE IF NOT EXISTS order_items (
--     id SERIAL PRIMARY KEY,
--     order_id INT REFERENCES orders(id),
--     product_id INT REFERENCES products(id),
--     quantity INT NOT NULL,
--     price DECIMAL(10, 2) NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS cart (
--     id SERIAL PRIMARY KEY,
--     user_id INT REFERENCES users(id) ON DELETE CASCADE
-- );

-- CREATE TABLE IF NOT EXISTS cart_items (
--     id SERIAL PRIMARY KEY,
--     cart_id INT REFERENCES cart(id),
--     product_id INT REFERENCES products(id),
--     quantity INT NOT NULL
-- );