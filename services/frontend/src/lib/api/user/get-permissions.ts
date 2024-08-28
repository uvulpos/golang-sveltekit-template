import ky from "ky";
import type { UserPermissions } from "./models/UserPermissions";
import { env } from "$env/dynamic/private";

export async function getUserPermissions(jwt: string): Promise<string[]> {
    return (await ky.get(env.BACKEND_URL + 'api/v1/self/permissions', { headers: { "Authorization": jwt } }).json<UserPermissions>()).permissions;
}