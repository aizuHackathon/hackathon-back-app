CREATE DATABASE test;

\c test

CREATE TABLE Users (
    ID SERIAL NOT NULL,
    Name VARCHAR(200),
    Height REAL,
    Weight REAL,
    Sex INTEGER,
    Old INTEGER,
    Created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (ID)
);

INSERT INTO Users VALUES (1, 'Aizu Taro', 164.2, 58.8, 0, 21),
    (2, 'Aizu Hanako', 164.2, 58.8, 0, 21);