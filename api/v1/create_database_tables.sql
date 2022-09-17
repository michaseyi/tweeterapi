-- create database tables

-- user table
CREATE TABLE IF NOT EXISTS users (
   `id` VARCHAR(60),
   `email` VARCHAR(60) NOT NULL UNIQUE,
   `username` VARCHAR(50) NOT NULL UNIQUE,
   `password` VARCHAR(128),
   PRIMARY KEY(`id`)
);

