-- +migrate Up

CREATE TABLE leagues (
  id SERIAL PRIMARY KEY,
  unique_name TEXT UNIQUE,
  display_name TEXT NOT NULL
);

CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  unique_name TEXT UNIQUE,
  display_name TEXT NOT NULL
);

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  league_id INT NOT NULL REFERENCES leagues(id),
  reporter_player_id INT NOT NULL REFERENCES players(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE details (
  player_id INT NOT NULL REFERENCES players(id),
  match_id INT NOT NULL REFERENCES matches(id),
  rating_change INT NOT NULL,
  has_won BOOLEAN NOT NULL
);

CREATE TABLE ratings (
  player_id INT NOT NULL REFERENCES players(id),
  league_id INT NOT NULL REFERENCES leagues(id),
  rating INT NOT NULL,
  CONSTRAINT ratings_pk PRIMARY KEY (player_id, league_id)
);

-- +migrate Down

DROP TABLE ratings;
DROP TABLE details;
DROP TABLE matches;
DROP TABLE players;
DROP TABLE leagues;
