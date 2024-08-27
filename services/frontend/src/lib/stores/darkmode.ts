import { writable, type Writable } from "svelte/store";
import Cookies from 'js-cookie'

const themeCookieName = "theme-style"
const THEME_COOKIE_DARKMODE = "darkmode"
const THEME_COOKIE_LIGHTMODE = "lightmode"

export enum ThemeEnum {
    Lightmode = 1,
    Darkmode
}

class ThemeStore {

    public themeStore: Writable<ThemeEnum | undefined> = writable<ThemeEnum | undefined>(undefined);

    constructor() {
        const themeCookie = Cookies.get(themeCookieName)

        if (themeCookie == THEME_COOKIE_DARKMODE) {
            this.toDarkMode()
        }
        else if (themeCookie == THEME_COOKIE_LIGHTMODE) {
            this.toLightMode()
        }
    }

    async toggleTheme() {
        const themeCookie = Cookies.get(themeCookieName)

        if (themeCookie === THEME_COOKIE_DARKMODE)
            this.toLightMode()
        else if (themeCookie === THEME_COOKIE_LIGHTMODE)
            this.toDarkMode()
    }

    isThemeSetOrAutoDetect(window: any) {
        const themeCookie = Cookies.get(themeCookieName)

        if (themeCookie == "" || themeCookie == null) {
            if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
                themeStore.toDarkMode()
            }
            else {
                themeStore.toLightMode()
            }
        }
    }

    toLightMode() {
        Cookies.set(themeCookieName, THEME_COOKIE_LIGHTMODE, { sameSite: "lax" })
        this.themeStore.set(ThemeEnum.Lightmode)
    }

    toDarkMode() {
        Cookies.set(themeCookieName, THEME_COOKIE_DARKMODE, { sameSite: "lax" })
        this.themeStore.set(ThemeEnum.Darkmode)
    }

    subscribe(run: any) {
        return this.themeStore.subscribe(run);
    }

}

export const themeStore = new ThemeStore();
