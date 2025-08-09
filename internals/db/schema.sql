CREATE TABLE IF NOT EXISTS author (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO author (name) VALUES('eko');
INSERT INTO author (name) VALUES('eko Widodo');


CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    completed BOOLEAN
);


INSERT INTO todos (title, completed) VALUES('Belajar Go', true);
INSERT INTO todos (title, completed) VALUES('Belajar JS', false);
INSERT INTO todos (title, completed) VALUES('Belajar PHP', false);

