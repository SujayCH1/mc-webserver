CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    discord_id VARCHAR(50) UNIQUE,
    discord_username VARCHAR(100),
    password_hash TEXT,
    whitelisted BOOLEAN DEFAULT false,
    banned BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE whitelist_requests (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    discord_id VARCHAR(50),
    discord_username VARCHAR(100),
    message TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE admin_users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    password_hash TEXT,
    role VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);