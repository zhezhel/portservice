--- Add ports table
CREATE TABLE portservice.ports (
    id      TEXT PRIMARY KEY,
    data    JSONB NOT NULL
);
--- apply above / revert below ---
DROP TABLE portservice.ports;
