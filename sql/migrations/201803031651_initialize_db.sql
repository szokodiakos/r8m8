-- +migrate Up

CREATE TABLE leagues (
  id TEXT PRIMARY KEY,
  display_name TEXT NOT NULL
);

CREATE TABLE players (
  id TEXT PRIMARY KEY,
  display_name TEXT NOT NULL
);

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  league_id TEXT NOT NULL REFERENCES leagues(id) ON DELETE CASCADE,
  reporter_player_id TEXT NOT NULL REFERENCES players(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE league_players (
  player_id TEXT NOT NULL REFERENCES players(id),
  league_id TEXT NOT NULL REFERENCES leagues(id) ON DELETE CASCADE,
  rating INT NOT NULL,
  PRIMARY KEY (player_id, league_id)
);

CREATE TABLE match_players (
  player_id TEXT NOT NULL,
  league_id TEXT NOT NULL,
  match_id INT NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
  rating_change INT NOT NULL,
  has_won BOOLEAN NOT NULL,
  FOREIGN KEY (player_id, league_id) REFERENCES league_players(player_id, league_id),
  PRIMARY KEY (player_id, league_id, match_id)
);

-- +migrate Down

DROP TABLE match_players;
DROP TABLE league_players;
DROP TABLE matches;
DROP TABLE players;
DROP TABLE leagues;
