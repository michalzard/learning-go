-- +goose Up
-- +goose StatementBegin
CREATE TABLE kv(
    k VARCHAR(255) NOT NULL,
    v TEXT NOT NULL, 
    created_at TIMESTAMP NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE kv;
-- +goose StatementEnd
