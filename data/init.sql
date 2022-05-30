CREATE TABLE IF NOT EXISTS settings
(
    id   SERIAL NOT NULL PRIMARY KEY,
    data jsonb
);