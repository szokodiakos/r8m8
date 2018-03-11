-- +migrate Up

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  unique_name TEXT UNIQUE,
  display_name TEXT NOT NULL,
  rating INT DEFAULT 1500
);

CREATE TABLE match_details (
  player_id INT NOT NULL REFERENCES players(id),
  match_id INT NOT NULL REFERENCES matches(id),
  rating_change INT NOT NULL
);

-- +migrate Down

DROP TABLE match_details;
DROP TABLE players;
DROP TABLE matches;
