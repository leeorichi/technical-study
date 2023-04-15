-- database: /home/longha/coding/technical-study/grapql/php-graphql/db/graphql.sqlite

-- Use the ▷ button in the top right corner to run the entire file.

CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT,
  email TEXT
);


INSERT INTO users (name, email) VALUES ('John Doe', 'john@example.com');
INSERT INTO users (name, email) VALUES ('Jane Doe 3', 'jane@example.com');


CREATE TABLE books (
  id INTEGER PRIMARY KEY,
  title TEXT,
  author TEXT
);

INSERT INTO books (title, author) VALUES ('Book: life of John Doe 1', 'John Doe');
INSERT INTO books (title, author) VALUES ('Book: life of John Doe 2', 'John Doe');
INSERT INTO books (title, author) VALUES ('Book: life of John Doe 3', 'John Doe');
INSERT INTO books (title, author) VALUES ('Book: life of John Doe 4', 'John Doe');
INSERT INTO books (title, author) VALUES ('Book: life of John Doe 5', 'John Doe');
INSERT INTO books (title, author) VALUES ('Book: life of John Doe 6', 'John Doe');
