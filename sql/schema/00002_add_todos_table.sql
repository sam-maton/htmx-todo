-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  completed BOOLEAN NOT NULL,
  description TEXT,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user_id FOREIGN KEY (user_id)
  REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
