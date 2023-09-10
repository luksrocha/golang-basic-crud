CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS houses (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR not null,
  address VARCHAR NOT NULL,
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default now()
);

CREATE TABLE IF NOT EXISTS users (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  first_name VARCHAR not null,
  last_name VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  hashed_password VARCHAR NOT NULL,
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default now()
);