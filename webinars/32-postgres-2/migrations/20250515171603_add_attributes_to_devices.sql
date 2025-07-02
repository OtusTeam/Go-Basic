-- +goose Up
-- +goose StatementBegin
alter table devices add column attributes jsonb;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table devices drop column attributes;
-- +goose StatementEnd
