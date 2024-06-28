CREATE TABLE "user" (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(128),
    email VARCHAR(128),
    date_birth DATE,
    registration_date DATE
);