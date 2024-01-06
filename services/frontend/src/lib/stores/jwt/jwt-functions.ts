import { jwtStore } from "./jwt";
import { goto } from "$app/navigation";
import { loginUserByCredentials } from "$lib/api/login";
import { refreshUserLoginJWT } from "$lib/api/login/update-jwt";

export async function loginUser(username: string, password: string): Promise<boolean> {
    const tokenString: string | null = await loginUserByCredentials(username, password)
    if (tokenString === null) {
        return false
    }

    jwtStore.set({
        jwtToken: tokenString,
    })

    goto("/dashboard")
    return true
}

export async function refreshUserLogin(username: string): Promise<boolean> {
    const tokenString: string | null = await refreshUserLoginJWT(username)
    if (tokenString === null) {
        return false
    }

    jwtStore.set({
        jwtToken: tokenString,
    })

    return true
}
