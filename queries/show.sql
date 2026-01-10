-- name: GetShowsByMovieAndCity :many
SELECT t.id AS theater_id,
       t.name AS theater_name,
       t.address,
       t.location,

       s.id AS show_id,
       s.start_time,

       f.id AS format_id,
       f.name AS format_name
FROM shows s
         JOIN screens sc ON sc.id = s.screen_id
         JOIN theaters t ON t.id = sc.theater_id
         JOIN formats f ON f.id = s.format_id
WHERE s.movie_id = $1
  AND t.city_code = $2
  AND s.start_time > now()
  AND s.status != 'cancelled'
ORDER BY
    t.name,
    s.start_time;
