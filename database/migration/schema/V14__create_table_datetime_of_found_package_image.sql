CREATE TABLE datetime_of_found_package_image (
    game_id INT UNSIGNED NOT NULL,
    executed_search_for_get_image_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (executed_search_for_get_image_id) REFERENCES executed_search_for_get_image(id),
    UNIQUE KEY (game_id, executed_search_for_get_image_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;