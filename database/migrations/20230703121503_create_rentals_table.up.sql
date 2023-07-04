CREATE TABLE rentals(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id INT(10) UNSIGNED NOT NULL,
    book_id INT(10) UNSIGNED NOT NULL,
    status ENUM('borrowed', 'returned','late') NOT NULL DEFAULT 'borrowed',

    rent_date DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    return_date DATE NULL DEFAULT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (book_id) REFERENCES books (id)
)