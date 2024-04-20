CREATE TABLE user_login (
  id SERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  total_login SERIAL NOT NULL,
  last_login TIMESTAMP,
  created_at TIMESTAMP,
  created_by BIGSERIAL
);

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  phone_no VARCHAR ( 13 ) NOT NULL,
  fullname VARCHAR ( 100 ) NOT NULL,
  password BYTEA NOT NULL,
  created_at TIMESTAMP,
  created_by BIGSERIAL,
  updated_at TIMESTAMP,
  updated_by BIGSERIAL
);