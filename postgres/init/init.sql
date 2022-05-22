CREATE DATABASE test;

\c test

CREATE TABLE users (
    ID SERIAL,
    Name VARCHAR(200),
    Height REAL,
    Weight REAL,
    Sex INTEGER,
    Old INTEGER,
    Created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (ID)
);

INSERT INTO users(Name, Height, Weight, Sex, Old) VALUES ('Aizu Taro', 164.2, 58.8, 0, 21),
    ('Aizu Hanako', 164.2, 58.8, 0, 21);

CREATE TABLE weights (
    ID SERIAL,
    User_ID SERIAL NOT NULL,
    Value REAL,
    Created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (ID),
    CONSTRAINT fk_Users FOREIGN KEY(User_ID) REFERENCES Users(ID) ON DELETE CASCADE
);

INSERT INTO Weights(User_ID, Value) VALUES (1, 78.2),
    (2, 58.8),
    (2, 45.8);

CREATE TABLE mongons (
    ID SERIAL,
    Mongon VARCHAR(200),
    PRIMARY KEY (ID)
);

INSERT INTO mongons VALUES (1, '運動しようぜ'),
    (2, '運動したいぜ'),
    (3, '運動しない？');
