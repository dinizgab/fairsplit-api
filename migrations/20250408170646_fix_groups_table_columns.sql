-- +goose Up
-- +goose StatementBegin
ALTER TABLE groups DROP COLUMN description;
ALTER TABLE groups ADD COLUMN value TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE groups DROP COLUMN value;
ALTER TABLE groups ADD COLUMN description TEXT;
-- +goose StatementEnd
