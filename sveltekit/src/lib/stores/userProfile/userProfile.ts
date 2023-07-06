import { writable } from "svelte/store";

export type userProfileType = {
    uuid: string
    username: string
    permisssions: string[]
}

export const userProfileStore = writable<userProfileType | undefined>(undefined)

