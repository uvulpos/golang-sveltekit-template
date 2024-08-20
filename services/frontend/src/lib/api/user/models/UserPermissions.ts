export interface UserPermissions {
    username: string;
    profilepicture: string;
    permissions: string[];
}

export const USER_PERMISSION_ADMIN = "admin";