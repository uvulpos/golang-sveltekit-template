import { getLocaleFromNavigator } from "svelte-i18n";
import Cookies from 'js-cookie'

export function getLocale(allowedLanguages: string[] = []): null | string {
    let language = null

    // read from browser
    let navigatorLanguage = getLocaleFromNavigator()
    if (navigatorLanguage !== null && navigatorLanguage !== undefined && allowedLanguages.includes(navigatorLanguage)) {
        language = navigatorLanguage
    }

    // read from cookie
    let cookieLanguage = Cookies.get("app-language")
    if (cookieLanguage !== undefined && allowedLanguages.includes(cookieLanguage)) {
        language = cookieLanguage
    }

    // cookie is the location where language should be stored
    if (language !== null) {
        Cookies.set("app-language", language)
    }

    console.log("Language result:", language);
    console.log("Language result:", navigatorLanguage);
    console.log("Language result:", cookieLanguage);

    return language
}