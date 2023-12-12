import { writable } from "svelte/store";
import Cookies from 'js-cookie'

export type jwtStoreType = {
    jwtToken: string,
    authCreated: Date,
    authRefreshed: Date,
    authExpires: Date,
}

export let jwtStore = writable<jwtStoreType | undefined>(undefined);

export type jwtDataType = {
    username: string;
}

export let jwtDataStore = writable<jwtDataType | undefined>(undefined);


// store as cookie
jwtStore.subscribe((jwt) => {
    if (jwt === undefined) {
        return
    }

    jwtDataStore.set({
        username: "uvulpos",
    })

    const jwtToken = JSON.stringify({
        jwtToken: jwt.jwtToken,
        authCreated: jwt.authCreated,
        authRefreshed: jwt.authRefreshed,
        authExpires: jwt.authExpires
    })

    Cookies.set("jwt", jwtToken, {
        expires: 30,
        domain: undefined,
        path: "/",
        secure: false,
        sameSite: undefined,
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
    jwtStore.set(undefined)
    jwtDataStore.set(undefined)
    Cookies.remove("jwt")
}