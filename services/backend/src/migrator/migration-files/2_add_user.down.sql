ALTER TABLE user_sessions DROP CONSTRAINT IF EXISTS created_ip_addr_not_empty_string;
ALTER TABLE user_sessions DROP CONSTRAINT IF EXISTS last_jwt_refresh_ip_addr_not_empty_string;

DROP TABLE IF EXISTS user_settings;
DROP TABLE IF EXISTS user_sessions;
DROP TABLE IF EXISTS user_identities;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS identity_provider;
