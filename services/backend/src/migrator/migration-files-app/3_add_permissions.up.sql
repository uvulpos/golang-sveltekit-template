CREATE TABLE IF NOT EXISTS roles (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    inherit_from UUID,
    FOREIGN KEY (inherit_from) REFERENCES roles(id) ON DELETE SET NULL,
    CHECK (inherit_from != id)
);

CREATE TABLE IF NOT EXISTS permissions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    identifier VARCHAR(255) UNIQUE NOT NULL
    CHECK (identifier != '')
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id UUID,
    permission_id UUID,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID,
    role_id UUID,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE OR REPLACE FUNCTION get_user_permissions(p_user_id UUID) 
RETURNS JSONB AS $$
BEGIN
    RETURN (
        WITH RECURSIVE role_hierarchy AS (
            -- Base case: Get all roles directly assigned to the user
            SELECT
                ur.user_id,
                r.id AS role_id
            FROM    user_roles ur
            JOIN    roles r ON ur.role_id = r.id
            WHERE   ur.user_id = p_user_id
            UNION
            
            -- Recursive case: Get all roles that are inherited from other roles
            SELECT
                rh.user_id,
                r2.id AS role_id
            FROM    role_hierarchy rh
            JOIN    roles r1 ON rh.role_id = r1.id
            JOIN    roles r2 ON r1.inherit_from = r2.id
        )
        SELECT 
            jsonb_agg(DISTINCT p.identifier) AS permissions
        FROM    role_hierarchy rh
        JOIN    role_permissions rp ON rh.role_id = rp.role_id
        JOIN    permissions p ON rp.permission_id = p.id
    );
END;
$$ LANGUAGE plpgsql;