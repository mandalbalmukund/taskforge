CREATE TABLE workflows (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    version INT NOT NULL,
    description TEXT,
    definition JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL
);