import type { SelfUser } from "./getUser"

export async function updateUser(username: string, email: string): Promise<SelfUser | null> {
    try {
        let request: Response
        let statusCode: number
        request = await fetch("/api/v1/self/update-user-data", {
            method: "POST",
            body: JSON.stringify({
                email,
                username,
            }),
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