-- +goose Up
-- +goose StatementBegin
CREATE TABLE films(
    ID int PRIMARY KEY NOT NULL ,
    Name TEXT NOT NULL DEFAULT '',
    Description TEXT NOT NULL DEFAULT '',
    Rating int NOT NULL DEFAULT 0,
    ReleaseDate date NOT NULL DEFAULT '2022-12-25'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE films;
-- +goose StatementEnd
