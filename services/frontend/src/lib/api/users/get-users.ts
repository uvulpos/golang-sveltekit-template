export type User = {
    "uuid": string;
    "username": string;
    "email": string;
    "auth_source": string
    "suspendet": boolean
    "role_name": string
}

export async function loginUserByCredentials(): Promise<User[] | null> {
    try {
        let request: Response
        let statusCode: number
        let responseText: string
        request = await fetch("/api/v1/users", {
            method: "GET",
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            console.log("status", statusCode);
            return null
        }

        return await awaitedResponse.json() as User[]

    } catch (error) {
        console.log(error);
        return null
    }
}