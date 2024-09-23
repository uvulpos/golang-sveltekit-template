import ky from "ky";

export async function refreshJwtToken(): Promise<string | null> {
    try {
        return (await ky.get('/api/v1/auth/refresh', { credentials: "include" }).json<string>());
    } catch (e) {
        return null
    }
}