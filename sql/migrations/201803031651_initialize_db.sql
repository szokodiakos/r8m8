-- +migrate Up

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  created_at timestamp NOT NULL
);

CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  rating int DEFAULT 1500
);

CREATE TABLE match_details (
  player_id int NOT NULL REFERENCES players(id),
  match_id int NOT NULL REFERENCES matches(id),
  rating_change int NOT NULL
);

CREATE TABLE slack_players (
  player_id int NOT NULL REFERENCES players(id),
  user_id varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  team_id varchar(255) NOT NULL
);

CREATE INDEX slack_players_user_id_team_id
ON slack_players (user_id, team_id);

-- +migrate Down

DROP TABLE slack_players;
DROP TABLE match_details;
DROP TABLE players;
DROP TABLE matches;
