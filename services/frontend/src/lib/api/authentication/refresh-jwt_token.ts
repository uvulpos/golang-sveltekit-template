import ky from "ky";

export async function refreshJwtToken(): Promise<string | null> {
    return (await ky.post('/api/v1/auth/refresh', { credentials: "include" }).json<string>());
}