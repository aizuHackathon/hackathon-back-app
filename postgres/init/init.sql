CREATE DATABASE test;

\c test

CREATE TABLE users (
    id SERIAL,
    name VARCHAR(200),
    height REAL,
    weight REAL,
    sex INTEGER,
    old INTEGER,
    created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (id)
);

INSERT INTO users(name, height, weight, sex, old) VALUES ('Aizu Taro', 164.2, 58.8, 0, 21),
    ('Aizu Hanako', 164.2, 58.8, 0, 21);

CREATE TABLE weights (
    id SERIAL,
    user_id SERIAL NOT NULL,
    value REAL,
    created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (id),
    CONSTRAINT fk_Users FOREIGN KEY(user_id) REFERENCES Users(id) ON DELETE CASCADE
);

INSERT INTO weights(user_id, value) VALUES (1, 78.2),
    (2, 58.8),
    (2, 45.8);