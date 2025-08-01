CREATE TABLE game_records (
    id SERIAL PRIMARY KEY,
    session_id TEXT,
    records JSONB,
    created_at TIMESTAMP DEFAULT now()
);
