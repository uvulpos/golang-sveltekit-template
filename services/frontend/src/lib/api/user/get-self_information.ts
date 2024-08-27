import type { SelfInformation } from "./models/SelfInformation";

export async function getSelfInformation(fetch: (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>, jwt: string): Promise<SelfInformation | null> {
    let response;

    try {
        const res = await fetch("http://backend:3000/api/v1/self", {
            method: "GET",
            headers: {
                Authorization: `${jwt}`,
                'Content-Type': 'application/json'
            },
        });

        if (!res.ok) {
            throw new Error('Network response was not ok');
        }

        response = await res.json() as SelfInformation;
    } catch (error) {
        return null;
    }

    return response;
}