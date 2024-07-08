import { getLocaleFromNavigator } from "svelte-i18n";
import Cookies from 'js-cookie'

const AppLanguageCookieName = "app-language"

export function setLocale(setLanguage: string, allowedLanguages: string[] = []) {
    if (allowedLanguages.includes(setLanguage)) {
        Cookies.set(AppLanguageCookieName, setLanguage)
    }
}

export function getLocale(allowedLanguages: string[] = []): null | string {
    let language = null

    // read from browser
    let navigatorLanguage = getLocaleFromNavigator()
    if (navigatorLanguage !== null && navigatorLanguage !== undefined && allowedLanguages.includes(navigatorLanguage)) {
        language = navigatorLanguage
    }

    // read from cookie
    let cookieLanguage = Cookies.get(AppLanguageCookieName)
    if (cookieLanguage !== undefined && allowedLanguages.includes(cookieLanguage)) {
        language = cookieLanguage
    }

    // cookie is the location where language should be stored
    if (language !== null) {
        Cookies.set(AppLanguageCookieName, language)
    }

    return language
}