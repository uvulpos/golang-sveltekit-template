import { setCookie as setCookieTS } from 'typescript-cookie'

export function setCookie(name: string, value: string, exdays: number, secure: boolean) {
    setCookieTS(
        name,
        value,
        {
            secure,
            sameSite: secure ? "strict" : "none",
            expires: exdays
        }
    )
}