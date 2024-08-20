import ky from "ky";
import type { UserPermissions } from "./models/UserPermissions";

export async function getUserPermissions(jwt: string): Promise<string[]> {
    return (await ky.get('http://backend:3000/api/v1/self/permissions', { headers: { "Authorization": jwt } }).json<UserPermissions>()).permissions;
}