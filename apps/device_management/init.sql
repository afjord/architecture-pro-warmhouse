-- Create the database if it doesn't exist
CREATE DATABASE device_management;

-- Connect to the database
\c device_management;

CREATE TABLE IF NOT EXISTS devices (
    id SERIAL PRIMARY KEY,
    home_id UUID, -- This field is nullable for simplicity.
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    location VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL,
    value double precision,
    unit VARCHAR(20),
    last_updated TIMESTAMP WITH TIME ZONE not null,
    created_at TIMESTAMP WITH TIME ZONE not null
);