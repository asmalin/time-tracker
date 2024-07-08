CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_number VARCHAR(50) UNIQUE NOT NULL,
    surname VARCHAR(100),
    name VARCHAR(100),
    patronymic VARCHAR(100),
    address VARCHAR(255)
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    name VARCHAR(100),
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE
);