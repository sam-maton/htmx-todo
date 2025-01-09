-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos DROP COLUMN description;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos ADD COLUMN description TEXT;
-- +goose StatementEnd
