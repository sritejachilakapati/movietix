-- name: GetMoviesByCity :many
SELECT DISTINCT
    m.*
FROM movies m
         JOIN shows s ON s.movie_id = m.id
         JOIN screens sc ON sc.id = s.screen_id
         JOIN theaters t ON t.id = sc.theater_id
WHERE t.city_code = $1
  AND s.start_time > now()
  AND s.status != 'cancelled'
ORDER BY m.rating DESC NULLS LAST
LIMIT $2
OFFSET $3;
