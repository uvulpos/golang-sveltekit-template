import { getUserPermissions } from "$lib/api/user/get-permissions";

export async function checkUserPermissionMiddleware(jwt: string, permissions: string[]): Promise<boolean> {
    const userPermissions = await getUserPermissions(jwt)

    for (const permission of permissions) {
        if (!userPermissions.includes(permission)) {
            return false;
        }
    }
    return true
}