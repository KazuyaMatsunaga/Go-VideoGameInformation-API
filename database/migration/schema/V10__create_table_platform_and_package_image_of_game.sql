CREATE TABLE platform_and_package_image_of_game (
    game_id INT UNSIGNED NOT NULL,
    platform_id INT UNSIGNED NOT NULL,
    package_image_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (platform_id) REFERENCES platform(id),
    FOREIGN KEY (package_image_id) REFERENCES package_image(id),
    UNIQUE KEY (game_id, platform_id, package_image_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;