CREATE TYPE identity_provider AS ENUM ();
ALTER TYPE identity_provider ADD VALUE 'Authentik';

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL, 
    email VARCHAR(255) NOT NULL, 
    email_verified VARCHAR(255) NOT NULL
);

CREATE TABLE user_identities (
    provider identity_provider NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (provider, provider_user_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (provider, provider_user_id, user_id)
);

CREATE TABLE user_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    useragent_hash VARCHAR(255) CHECK (useragent_hash != ''), -- to prevent at least a bit session hijacking
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    created_ip_addr VARCHAR(255) DEFAULT NULL,
    last_jwt_refresh TIMESTAMP NOT NULL DEFAULT NOW(),
    last_jwt_refresh_ip_addr VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (id, user_id)
);