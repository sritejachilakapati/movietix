CREATE TABLE cities (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

ALTER TABLE theaters ADD COLUMN city_code TEXT REFERENCES cities(code) NOT NULL;