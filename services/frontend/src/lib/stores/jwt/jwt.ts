import { writable } from "svelte/store";
import Cookies from 'js-cookie'
import { parseJWT } from "$lib/functions/parse-jwt";

export type jwtStoreType = {
    jwtToken: string,
    authCreated: Date,
    authRefreshed: Date,
    authExpires: Date,
}

export type jwtDataType = {
    username: string;
}

export let jwtStore = writable<jwtStoreType | undefined>(undefined);
export let jwtDataStore = writable<jwtDataType | undefined>(undefined);

// store as cookie
jwtStore.subscribe((jwt) => {
    if (jwt === undefined) {
        jwtDataStore.set(undefined);
        return
    }

    // process answer
    const parsedJWT = parseJWT(jwt.jwtToken)
    jwtDataStore.set({
        // uuid: parsedJWT?.{ "user-uuid"},
        username: parsedJWT?.username,
        // roles: parsedJWT?.roles,
    })

    const jwtToken = JSON.stringify({
        jwtToken: jwt.jwtToken,
        authCreated: jwt.authCreated,
        authRefreshed: jwt.authRefreshed,
        authExpires: jwt.authExpires
    })

    Cookies.set("jwt", jwtToken, {
        sameSite: "lax",
    })
})

export function loginUserFromCookie() {
    const userJWTString = Cookies.get("jwt")
    if (userJWTString === undefined) {
        return
    }

    const userJWT = JSON.parse(userJWTString) as jwtStoreType
    jwtStore.set({
        jwtToken: userJWT.jwtToken,
        authCreated: userJWT.authCreated,
        authRefreshed: userJWT.authRefreshed,
        authExpires: userJWT.authExpires,
    })
}

export function jwtSessionDestroy() {
    jwtStore.set(undefined);
    Cookies.remove("jwt");
}