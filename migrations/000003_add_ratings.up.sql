ALTER TABLE movies
    ADD COLUMN certification TEXT,        -- PG-13, R, U/A
    ADD COLUMN rating_score NUMERIC(3,1);  -- 0.0 – 10.0
