CREATE TABLE genre_of_game (
    game_id INT UNSIGNED NOT NULL,
    genre_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (genre_id) REFERENCES genre(id),
    UNIQUE KEY (game_id, genre_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;