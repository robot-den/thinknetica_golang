-- выборка фильмов с названием студии
SELECT movies.name, pc.name as company
FROM movies
    JOIN studios pc on movies.studio_id = pc.id;

-- выборка фильмов для некоторого актёра
SELECT movies.name
FROM movies
    JOIN movies_actors ma on movies.id = ma.movie_id
    JOIN actors a on ma.actor_id = a.id
WHERE (a.first_name || ' ' || a.second_name) = 'Morgan Freeman';

-- подсчёт фильмов для некоторого режиссёра
SELECT count(*)
FROM movies
    JOIN movies_directors mp on movies.id = mp.movie_id
    JOIN directors p on p.id = mp.director_id
WHERE (p.first_name || ' ' || p.second_name) = 'Some Guy 2';

-- выборка фильмов для нескольких режиссёров из списка (подзапрос)
SELECT DISTINCT movies.name
FROM movies JOIN movies_directors mp on movies.id = mp.movie_id
WHERE mp.director_id IN (
    SELECT directors.id
    FROM directors
    WHERE (directors.first_name || ' ' || directors.second_name) IN ('Some Guy 2', 'Some Guy 4')
)
ORDER BY movies.name;

-- подсчёт количества фильмов для актёра
SELECT count(*)
FROM movies
         JOIN movies_actors ma on movies.id = ma.movie_id
         JOIN actors a on ma.actor_id = a.id
WHERE (a.first_name || ' ' || a.second_name) = 'Morgan Freeman';

-- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах
SELECT role, full_name, movies_count
FROM (
      SELECT
         'Actor' as role,
         (first_name || ' ' || second_name) as full_name,
         count(*) as movies_count
      FROM
           movies_actors JOIN actors a on a.id = movies_actors.actor_id
      GROUP BY (first_name || ' ' || second_name)
      HAVING count(*) > 1
    ) as t
UNION ALL
(
    SELECT
        'Director' as role,
        (first_name || ' ' || second_name) as full_name,
        count(*) as movies_count
    FROM
        movies_directors JOIN directors p on p.id = movies_directors.director_id
    GROUP BY (first_name || ' ' || second_name)
    HAVING count(*) > 1
)
ORDER BY movies_count DESC;

-- подсчёт количества фильмов со сборами больше 1000
SELECT count(*)
FROM movies
WHERE gross > 1000;

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000
SELECT count(DISTINCT director_id)
FROM movies_directors
WHERE movie_id IN (
    SELECT id
    FROM movies
    WHERE gross > 1000
);

-- выборка различных фамилий актёров
SELECT DISTINCT second_name
FROM actors
ORDER BY second_name;

-- подсчёт количества фильмов, имеющих дубли по названию
SELECT count(*)
FROM (
    SELECT name
    FROM movies
    GROUP BY name
    HAVING count(*) > 1
) AS t;


