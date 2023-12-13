import Cookies from "js-cookie";
import { jwtStore } from "./jwt";
import { goto } from "$app/navigation";

export function loginUser(username: string, password: string) {
    // request the backend

    // handle errors

    // process answer

    // set jwt
    let now: Date = new Date()
    let expires = now
    expires.setFullYear(2024)
    jwtStore.set({
        jwtToken: `${username} ${password}`,
        authExpires: expires,
        authCreated: now,
        authRefreshed: now
    })
    goto("/dashboard")
}
