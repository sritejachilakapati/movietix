-- name: GetShowsByMovieAndCity :many
SELECT m.id    AS movie_id,
       m.title AS movie_title,
       t.id    AS theater_id,
       t.name  AS theater_name,
       t.address,
       t.location,
       s.id    AS show_id,
       s.start_time,
       f.id    AS format_id,
       f.name  AS format_name
FROM theaters t
         JOIN screens sc ON sc.theater_id = t.id
         JOIN shows s ON s.screen_id = sc.id
         JOIN movies m on m.id = s.movie_id
         JOIN formats f ON f.id = s.format_id
WHERE t.city_code = $1
  AND s.movie_id = $2
  AND s.start_time > now()
  AND s.status != 'cancelled'
ORDER BY t.id, s.start_time
    LIMIT $3
OFFSET $4;
