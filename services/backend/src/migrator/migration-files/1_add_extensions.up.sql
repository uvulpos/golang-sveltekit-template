CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE identity_provider AS ENUM ('Authentik');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR UNIQUE NOT NULL,
    display_name VARCHAR NOT NULL, 
    email VARCHAR NOT NULL, 
    email_verified VARCHAR NOT NULL
);

CREATE TABLE user_identities (
    provider identity_provider NOT NULL,
    provider_user_id VARCHAR NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (provider, provider_user_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    UNIQUE (provider, provider_user_id, user_id)
);

CREATE TABLE user_sessions (
    id UUID DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    useragent_hash VARCHAR CHECK (useragent_hash <> ''), -- to prevent at least a bit session hijacking
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    created_ip_addr VARCHAR DEFAULT NULL,
    last_jwt_refresh TIMESTAMP NOT NULL DEFAULT NOW(),
    last_jwt_refresh_ip_addr VARCHAR DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    UNIQUE (id, user_id)
);

-- Value can be null, but not empty string
ALTER TABLE user_sessions
ADD CONSTRAINT created_ip_addr_not_empty_string
CHECK (created_ip_addr IS NULL OR (created_ip_addr <> ''));

-- Value can be null, but not empty string
ALTER TABLE user_sessions
ADD CONSTRAINT last_jwt_refresh_ip_addr_not_empty_string
CHECK (last_jwt_refresh_ip_addr IS NULL OR (last_jwt_refresh_ip_addr <> ''));