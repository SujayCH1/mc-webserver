CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    version TEXT,
    port INTEGER UNIQUE,
    status TEXT DEFAULT 'stopped'
);