CREATE TABLE platform_and_price_of_game (
    game_id INT UNSIGNED NOT NULL,
    platform_id INT UNSIGNED NOT NULL,
    price VARCHAR(255) NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (platform_id) REFERENCES platform(id),
    UNIQUE KEY (game_id, platform_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;