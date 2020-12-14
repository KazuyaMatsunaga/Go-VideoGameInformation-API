CREATE TABLE datetime_of_found_game (
    game_id INT UNSIGNED NOT NULL,
    executed_scraping_for_get_game_information_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (executed_scraping_for_get_game_information_id) REFERENCES executed_scraping_for_get_game_information(id),
    UNIQUE KEY (game_id, executed_scraping_for_get_game_information_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;