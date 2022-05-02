CREATE DATABASE test;

\c test

CREATE TABLE Mongons (
    ID SERIAL NOT NULL,
    Mongon VARCHAR(200),
    PRIMARY KEY (ID)
);

INSERT INTO Mongons VALUES (1, '運動しようぜ'),
    (2, '運動したいぜ');