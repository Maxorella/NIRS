CREATE TABLE user (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(255),
    email VARCHAR(255),
    date_birth DATE,
    registration_date DATE
);