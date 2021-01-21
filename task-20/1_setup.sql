-- Create database if needed. I worked right in `postgres` database
-- DROP DATABASE IF EXISTS imdb;
-- CREATE DATABASE imdb;

-- Drop tables if they are exist to have idempotent
DROP INDEX IF EXISTS movies_release_year_name_idx, movies_actors_uniq_idx, movies_directors_uniq_idx;
DROP TABLE IF EXISTS movies, actors, directors, studios, movies_actors, movies_directors;
DROP TYPE IF EXISTS rating;

-- In real application I wouldn't use enum in database (it may change and require migration)
CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');

CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    release_year INTEGER NOT NULL CHECK ( release_year >= 1800 ),
    rating RATING NOT NULL,
    gross BIGINT NOT NULL DEFAULT 0,
    studio_id INTEGER REFERENCES studios
);
CREATE UNIQUE INDEX movies_release_year_name_idx ON movies (release_year, name);

CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE movies_actors (
    movie_id INTEGER REFERENCES movies,
    actor_id INTEGER REFERENCES actors
);
CREATE UNIQUE INDEX movies_actors_uniq_idx ON movies_actors (movie_id, actor_id);


CREATE TABLE movies_directors (
    movie_id INTEGER REFERENCES movies,
    director_id INTEGER REFERENCES directors
);
CREATE UNIQUE INDEX movies_directors_uniq_idx ON movies_directors (movie_id, director_id);
