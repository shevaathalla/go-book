ALTER TABLE publishers
    ADD COLUMN address VARCHAR(255) NOT NULL AFTER name,
    ADD COLUMN phone VARCHAR(255) NOT NULL AFTER address;