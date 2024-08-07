CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE identity_provider AS ENUM ('Authentik');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL, 
    email VARCHAR(255) NOT NULL, 
    email_verified VARCHAR NOT NULL
);

CREATE TABLE user_identities (
    provider identity_provider NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (provider, provider_user_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    UNIQUE (provider, provider_user_id, user_id)
);