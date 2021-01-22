INSERT INTO studios (name) VALUES
('Undefined'),
('Company 1'),
('Company 2'),
('Company 3'),
('Company 4'),
('Company 5');

INSERT INTO movies (name, release_year, rating, gross, studio_id) VALUES
('Movie 1', '1994', 'PG-13', 2000, 1),
('Movie 2', '1999', 'PG-10', 500, 1),
('Movie 3', '2000', 'PG-13', 4000, 1),
('Movie 4', '2001', 'PG-18', 1000, 1),
('Movie 5', '1990', 'PG-10', 990, 1),
('Movie 6', '1945', 'PG-10', 1500, 1),
('Movie 7', '2019', 'PG-18', 2500, 1),
('Movie 7', '2020', 'PG-18', 200, 1);

WITH relations AS (
    SELECT * FROM (
                      VALUES ('Movie 1', 'Company 5'),
                             ('Movie 2', 'Company 3'),
                             ('Movie 3', 'Company 1'),
                             ('Movie 4', 'Company 1'),
                             ('Movie 5', 'Company 2'),
                             ('Movie 6', 'Company 4'),
                             ('Movie 7', 'Company 1')
                  ) AS t (movie_name, studio_name)
)
UPDATE movies AS mvs
SET studio_id = studios.id
    FROM (relations JOIN studios ON studios.name = studio_name)
WHERE mvs.name = movie_name;

INSERT INTO actors (first_name, second_name, birthday) VALUES
('Morgan', 'Freeman', '1937-06-01'),
('Tim', 'Robbins', '1958-10-16'),
('Jack', 'Nicolson', '1950-10-16'),
('Robert', 'Downey Jr.', '1965-10-16'),
('Julianne', 'Moore', '1960-12-03'),
('Jim', 'Carrey', '1962-01-17');

INSERT INTO directors (first_name, second_name, birthday) VALUES
('Some', 'Guy 1', '1959-01-28'),
('Some', 'Guy 2', '1960-01-28'),
('Some', 'Guy 3', '1964-01-28'),
('Some', 'Guy 4', '1983-01-28'),
('Some', 'Guy 5', '1999-01-28');


WITH relations AS (
    SELECT * FROM (
                      VALUES ('Morgan Freeman', 'Movie 1'),
                             ('Morgan Freeman', 'Movie 3'),
                             ('Morgan Freeman', 'Movie 5'),
                             ('Tim Robbins', 'Movie 5'),
                             ('Tim Robbins', 'Movie 7'),
                             ('Jack Nicolson', 'Movie 2'),
                             ('Jack Nicolson', 'Movie 4'),
                             ('Jack Nicolson', 'Movie 6'),
                             ('Robert Downey Jr.', 'Movie 4'),
                             ('Robert Downey Jr.', 'Movie 6'),
                             ('Robert Downey Jr.', 'Movie 7'),
                             ('Julianne Moore', 'Movie 2'),
                             ('Julianne Moore', 'Movie 7'),
                             ('Jim Carrey', 'Movie 1'),
                             ('Jim Carrey', 'Movie 3')
                  ) AS t (actor_name, movie_name)
)
INSERT INTO movies_actors (movie_id, actor_id)
SELECT movies.id, actors.id FROM relations
                                     JOIN movies ON movies.name = movie_name
                                     JOIN actors ON (actors.first_name || ' ' || actors.second_name) = actor_name;

WITH relations AS (
    SELECT * FROM (
                      VALUES ('Some Guy 1', 'Movie 7'),
                             ('Some Guy 2', 'Movie 1'),
                             ('Some Guy 2', 'Movie 3'),
                             ('Some Guy 2', 'Movie 6'),
                             ('Some Guy 3', 'Movie 1'),
                             ('Some Guy 4', 'Movie 3'),
                             ('Some Guy 4', 'Movie 4'),
                             ('Some Guy 4', 'Movie 5'),
                             ('Some Guy 5', 'Movie 2')
                  ) AS t (director_name, movie_name)
)
INSERT INTO movies_directors (movie_id, director_id)
SELECT movies.id, directors.id FROM relations
                                        JOIN movies ON movies.name = movie_name
                                        JOIN directors ON (directors.first_name || ' ' || directors.second_name) = director_name;