CREATE TABLE IF NOT EXISTS password_resets(
  email VARCHAR(100) PRIMARY KEY,
  token VARCHAR(150) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX idx_password_resets_token ON password_resets (token);