CREATE TABLE users(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,

    code_verified VARCHAR(255) NULL DEFAULT NULL,
    email_verified_at TIMESTAMP NULL DEFAULT NULL,

    instance VARCHAR(255) NULL DEFAULT NULL,
    address TEXT NULL DEFAULT NULL,
    phone VARCHAR(255) NULL DEFAULT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),
    UNIQUE INDEX users_email_unique (email),
    INDEX users_name_index (name),
    INDEX users_email_index (email)
)