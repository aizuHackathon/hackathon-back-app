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

CREATE TABLE Keihatu(
    ID SERIAL NOT NULL,
    Value VARCHAR(200)
);

CREATE TABLE Weights (
    ID SERIAL NOT NULL,
    User_ID SERIAL NOT NULL,
    Value REAL,
    Created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (ID),
    CONSTRAINT fk_Users FOREIGN KEY(User_ID) REFERENCES Users(ID) ON DELETE CASCADE
);

INSERT INTO Users(ID, Name, Height, Weight,Sex, Old) VALUES 
(1, 'Aizu Taro', 164.2, 58.8, 0, 21),
(2, 'Aizu Hanako', 164.2, 58.8, 0, 21);

INSERT INTO Keihatu(ID, Value) VALUES 
(0, 'Everyone wants to be successful until they see what it actually takes.'),
(1, 'Its difficult to beat a person who never gives up.'),
(2, 'It doesnt matter what your background is and where you come from, if you have dreams and goals, thats all that matters.'),
(3, 'Its not the will to win that matters. Everyone has that. Its the will to prepare to win that matters.'),
(4, 'Be quick, but dont hurry.');

INSERT INTO Weights(ID, User_ID, Value) VALUES 
(1, 1, 78.2),
(2, 2, 58.8),
(3, 2, 45.8);

