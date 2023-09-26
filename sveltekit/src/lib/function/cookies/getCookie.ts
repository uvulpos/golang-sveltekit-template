import { getCookie as getCookieTS } from 'typescript-cookie'

export function getCookie(name: string): string | undefined {
    return getCookieTS(name)
}