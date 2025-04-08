-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_groups (
	primary key (user_id, group_id),
	user_id uuid REFERENCES users(id) ON DELETE CASCADE,
	group_id uuid REFERENCES groups(id) ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE groups ADD COLUMN owner_id uuid REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_groups;
-- +goose StatementEnd
