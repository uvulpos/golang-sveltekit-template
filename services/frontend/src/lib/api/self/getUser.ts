export type UserPermission = {
    "name": string;
    "identifier": string;
}

export type SelfUser = {
    "uuid": string;
    "username": string;
    "email": string;
    "auth_source": string
    "suspendet": boolean
    "roleName": string
    "permissions": Array<UserPermission>
}

export async function getUser(): Promise<SelfUser | null> {
    try {
        let request: Response
        let statusCode: number

        request = await fetch("/api/v1/self/get-user-data", {
            method: "GET",
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return null
        }
        return await awaitedResponse.json() as SelfUser
    } catch (error) {
        return null
    }
}