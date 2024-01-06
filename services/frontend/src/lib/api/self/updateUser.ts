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
        statusCode = request.status
        console.log("responseCode", statusCode);


        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return null
        }

        const myValue = await request.json()
        console.log("inside api", { myValue }); // <- hier hat er alte Daten. Woher ????

        return myValue as SelfUser
    } catch (error) {
        return null
    }
}