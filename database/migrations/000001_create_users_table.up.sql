CREATE TABLE IF NOT EXISTS `users` (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name varchar(191) DEFAULT NULL,
    username varchar(191) NOT NULL,
    email varchar(191) NOT NULL,
    password varchar(191) NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT NULL
);