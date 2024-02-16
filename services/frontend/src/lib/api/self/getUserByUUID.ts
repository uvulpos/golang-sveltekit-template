import type { UserPermission } from "./getUser";

export type ManagedUser = {
    "uuid": string;
    "username": string;
    "email": string;
    "auth_source": string
    "suspendet": boolean
    "role_name": string
    "permissions": Array<UserPermission>
}

export async function getUserByUUID(uuid: string): Promise<ManagedUser | null> {
    try {
        let request: Response
        let statusCode: number

        request = await fetch(`/api/v1/user/${uuid}`, {
            method: "GET",
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return null
        }
        return await awaitedResponse.json() as ManagedUser
    } catch (error) {
        return null
    }
}