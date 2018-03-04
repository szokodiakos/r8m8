-- +migrate Up

ALTER DATABASE r8m8 CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE matches (
  id int NOT NULL AUTO_INCREMENT,
  created_at timestamp NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE players (
  id int NOT NULL AUTO_INCREMENT,
  rating int DEFAULT 1500,
  PRIMARY KEY (id)
);

CREATE TABLE match_details (
  player_id int NOT NULL,
  match_id int NOT NULL,
  rating_change int NOT NULL,
  FOREIGN KEY (player_id) REFERENCES players(id),
  FOREIGN KEY (match_id) REFERENCES matches(id)
);

CREATE TABLE slack_players (
  player_id int NOT NULL,
  user_id varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  team_id varchar(255) NOT NULL,
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE INDEX slack_players_user_id_team_id
ON slack_players (user_id, team_id);

-- +migrate Down

DROP TABLE slack_players;
DROP TABLE match_details;
DROP TABLE players;
DROP TABLE matches;
