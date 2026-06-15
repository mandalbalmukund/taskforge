CREATE TABLE approvals (
    id UUID PRIMARY KEY,
    execution_id UUID NOT NULL,
    task_id VARCHAR(255) NOT NULL,
    approver VARCHAR(255),
    decision VARCHAR(50),
    approved_at TIMESTAMP
);