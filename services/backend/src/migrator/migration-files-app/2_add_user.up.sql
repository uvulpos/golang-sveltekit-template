CREATE TYPE identity_provider IF NOT EXISTS AS ENUM ();
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

-- Value can be null, but not empty string
ALTER TABLE user_sessions
ADD CONSTRAINT created_ip_addr_not_empty_string
CHECK (created_ip_addr IS NOT NULL AND IS NOT EMPTY);

-- Value can be null, but not empty string
ALTER TABLE user_sessions
ADD CONSTRAINT last_jwt_refresh_ip_addr_not_empty_string
CHECK (last_jwt_refresh_ip_addr IS NOT NULL AND IS NOT EMPTY);

CREATE TYPE app_language IF NOT EXISTS AS ENUM ();
ALTER TYPE app_language ADD VALUE 'en';
ALTER TYPE app_language ADD VALUE 'de';

CREATE TYPE app_theme IF NOT EXISTS AS ENUM ();
ALTER TYPE app_theme ADD VALUE 'lightmode';
ALTER TYPE app_theme ADD VALUE 'darkmode';

CREATE TABLE user_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    app_language identity_provider NOT NULL DEFAULT 'en',
    app_theme app_theme NOT NULL DEFAULT 'lightmode',
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (id, user_id)
);