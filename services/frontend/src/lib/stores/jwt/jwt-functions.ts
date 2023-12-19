import { jwtStore } from "./jwt";
import { goto } from "$app/navigation";
import { loginUserByCredentials } from "$lib/api/login";

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
