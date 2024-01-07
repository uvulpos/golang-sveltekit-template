import { writable } from "svelte/store";
import Cookies from 'js-cookie'
import { parseJWT } from "$lib/functions/parse-jwt";

export type jwtStoreType = {
    jwtToken: string,
}

export type jwtDataType = {
    uuid: string;
    username: string;
    roleName: string;
    permissions: string[];
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
    console.log(parsedJWT);

    jwtDataStore.set({
        uuid: parsedJWT?.["user-uuid"],
        username: parsedJWT?.username,
        roleName: parsedJWT?.roleName,
        permissions: parsedJWT?.permissions,
    })

    Cookies.set("jwt", jwt.jwtToken, {
        sameSite: "lax",
    })
})

export function loginUserFromCookie() {
    const userJWTString = Cookies.get("jwt")
    if (userJWTString === undefined) {
        return
    }

    jwtStore.set({
        jwtToken: userJWTString,
    })
}

export function jwtSessionDestroy() {
    jwtStore.set(undefined);
    Cookies.remove("jwt");
}