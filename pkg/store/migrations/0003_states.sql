--- Add states table
CREATE TABLE portservice.states (
    filename    TEXT PRIMARY KEY,
    offset_     BIGINT NOT NULL DEFAULT 0
);
--- apply above / revert below ---
DROP TABLE portservice.states;
