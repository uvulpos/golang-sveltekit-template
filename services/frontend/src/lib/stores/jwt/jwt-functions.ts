import { jwtStore } from "./jwt";
import { goto } from "$app/navigation";

export async function loginUser(username: string, password: string): Promise<boolean> {
    try {
        // request the backend
        let request: Response
        let statusCode: number
        let responseText: string
        request = await fetch("/api/login", {
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            }),
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status
        responseText = await awaitedResponse.text()

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            // do error handling
            return false
        }

        // set jwt
        let now: Date = new Date()
        let expires = now
        expires.setFullYear(2024)
        jwtStore.set({
            jwtToken: responseText,
            authExpires: expires,
            authCreated: now,
            authRefreshed: now
        })
        goto("/dashboard")
        return true
    } catch (error) {
        return false
    }
}
