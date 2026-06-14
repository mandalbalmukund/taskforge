CREATE TABLE task_executions (
    id UUID PRIMARY KEY,
    execution_id UUID NOT NULL,
    task_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    attempts INT DEFAULT 0,
    output JSONB,
    error TEXT,
    CONSTRAINT fk_execution
        FOREIGN KEY(execution_id)
        REFERENCES executions(id)
);