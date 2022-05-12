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

INSERT INTO Users(ID, Name, Height, Weight, Sex, Old) VALUES (1, 'Aizu Taro', 164.2, 58.8, 0, 21),
    (2, 'Aizu Hanako', 164.2, 58.8, 0, 21);

CREATE TABLE Weights (
    ID SERIAL NOT NULL,
    User_ID SERIAL NOT NULL,
    Value REAL,
    Created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (ID),
    CONSTRAINT fk_Users FOREIGN KEY(User_ID) REFERENCES Users(ID) ON DELETE CASCADE
);

INSERT INTO Weights(ID, User_ID, Value) VALUES (1, 1, 78.2),
    (2, 2, 58.8),
    (3, 2, 45.8);