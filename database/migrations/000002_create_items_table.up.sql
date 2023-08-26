CREATE TABLE IF NOT EXISTS Items(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER DEFAULT NULL,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    description TEXT NOT NULL,
    image TEXT NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT NULL,
    FOREIGN KEY(user_id) REFERENCES Users(id)
);