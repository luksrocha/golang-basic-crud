CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS houses (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR not null,
  address VARCHAR NOT NULL,
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default now()
);