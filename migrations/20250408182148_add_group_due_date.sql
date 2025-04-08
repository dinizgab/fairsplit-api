-- +goose Up
-- +goose StatementBegin
ALTER TABLE groups ADD COLUMN due_day smallint not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE groups DROP COLUMN due_day;
-- +goose StatementEnd
