DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS todos;

CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  uuid        VARCHAR(64) NOT NULL UNIQUE,
  name        VARCHAR(255),
  email       VARCHAR(255) NOT NULL UNIQUE,
  password    VARCHAR(255) NOT NULL,
  created_at  TIMESTAMP NOT NULL
);

CREATE TABLE sessions (
  id          SERIAL PRIMARY KEY,
  uuid        VARCHAR(64) NOT NULL UNIQUE,
  name        VARCHAR(255),
  email       VARCHAR(255),
  user_id     INTEGER REFERENCES users(id),
  created_at  TIMESTAMP NOT NULL
);

CREATE TABLE todos (
  id          SERIAL PRIMARY KEY,
  content     TEXT,
  user_id     INTEGER REFERENCES users(id),
  created_at  TIMESTAMP NOT NULL
);
