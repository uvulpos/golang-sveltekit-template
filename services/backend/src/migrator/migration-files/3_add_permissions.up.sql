CREATE TABLE IF NOT EXISTS roles (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    inherit_from UUID,
    FOREIGN KEY (inherit_from) REFERENCES roles(id) ON DELETE SET NULL,
    CHECK (inherit_from <> id)
);

CREATE TABLE IF NOT EXISTS permissions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    identifier VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id UUID,
    permission_id UUID,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);
