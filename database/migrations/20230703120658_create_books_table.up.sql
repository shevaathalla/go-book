CREATE TABLE books (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    stock INT NOT NULL,

    author_id INT(10) UNSIGNED NOT NULL,
    publisher_id INT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (author_id) REFERENCES authors (id),
    FOREIGN KEY (publisher_id) REFERENCES publishers (id),

    INDEX books_title_index (title)
);