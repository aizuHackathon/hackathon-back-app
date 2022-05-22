CREATE DATABASE test;

\c test

CREATE TABLE users (
    id SERIAL,
    name VARCHAR(200),
    height REAL,
    sex INTEGER,
    old INTEGER,
    pass VARCHAR(200),
    state INTEGER DEFAULT 0,
    created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE keihatu(
    id SERIAL ,
    value VARCHAR(200)
);

CREATE TABLE weights (
    id SERIAL,
    user_id SERIAL NOT NULL,
    value REAL,
    created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (id),
    CONSTRAINT fk_Users FOREIGN KEY(user_id) REFERENCES Users(id) ON DELETE CASCADE
);

INSERT INTO users(name, height, sex, old, pass) VALUES 
('Aizu Taro', 164.2, 0, 21, 'test'),
('Aizu Hanako', 164.2, 0, 21, 'test');

INSERT INTO keihatu(value) VALUES 
('筋肉は裏切らない。'),
('好きな女の子を振り向かせられるのは筋肉だけさ、'),
('そろそろ僕の上腕二頭筋が進化しそうだ、'),
('マッスルマッスル'),
('君ならできるよ。'),
('今日も一日頑張ろう！'),
('腹筋１００回、腕立て１００回、スクワット１００回！');

INSERT INTO Weights(ID, User_ID, Value) VALUES 
(1, 1, 78.2),
(2, 2, 58.8),
(3, 2, 45.8);


