import { writable, type Writable } from "svelte/store";

export interface ToastInformation {
    id: number;
    header: string;
    body: string;
    type: ToastType;
}

export type ToastType = "success" | "error"

function createToast() {
    const { subscribe, update, set } = writable(new Array())
    let count: number = 0

    // kept for future use
    // function _init() {}

    function push(type: ToastType, header: string, msg: string) {
        const entry: ToastInformation = {
            id: ++count,
            header: header,
            body: msg,
            type: type,
        }

        update((oldStore) => {
            return [...oldStore, entry]
        })
        return count
    }

    function pop(id: number) {
        update((entries) => {
            if (!entries.length || id === 0) return []
            const found = id || Math.max(...entries.map((i) => i.id))
            return entries.filter((i) => i.id !== found)
        })
    }

    return { subscribe, set, update, push, pop }
}

export const toast = createToast()