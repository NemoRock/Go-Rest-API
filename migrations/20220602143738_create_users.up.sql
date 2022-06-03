CREATE TABLE users (
    id bigserial not null primaty key,
    email vachar not null unique,
    encrypted_password varchar not null
)