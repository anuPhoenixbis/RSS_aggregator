-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea),'hex')
);
-- here we are in random()::text::bytea , generating a random slice of bytes, then we are hashing it with sha256 and encoding it to hexadecimal value
-- sha256 is a cryptographic hash function that generates a 256-bit (32-byte) hash value


-- down migration should undo what ever was done in up migration
-- +goose Down
ALTER TABLE users DROP COLUMN api_key;