-- +goose Up
-- +goose StatementBegin
CREATE TABLE films(
    ID int PRIMARY KEY NOT NULL ,
    Name TEXT NOT NULL DEFAULT '',
    Description TEXT NOT NULL,
    Rating int NOT NULL,
    ReleaseDate date NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE films;
-- +goose StatementEnd
