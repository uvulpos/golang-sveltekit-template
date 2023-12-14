import { jwtStore } from "./jwt";
import { goto } from "$app/navigation";

export async function loginUser(username: string, password: string) {
    // request the backend

    const request = await fetch("http://web.localhost/api/login", {
        method: "POST",
        body: JSON.stringify({
            username: username,
            password: password
        }),
        headers: new Headers({ 'content-type': 'application/json' }),
    })

    // handle errors
    const awaitedResponse = await request
    const statusCode = awaitedResponse.status
    const responseText = await awaitedResponse.text()
    if (!request.ok || statusCode < 200 || statusCode > 299) {
        // do error handling
        console.log("error, statuscode:", statusCode);
        console.log(responseText);
        return
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
}
