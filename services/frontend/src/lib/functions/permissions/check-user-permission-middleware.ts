import { checkUserPermission } from "./check-user-permission";

export async function checkUserPermissionMiddleware(fetch: (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>, jwt: string, permissions: string[]): Promise<boolean> {
    return await checkUserPermission(fetch, jwt, permissions);
}

export class NotPermittedError extends Error {
    constructor() {
        const message = "User not Authorized"
        super(message);
        this.name = 'NotPermittedError';
        this.message = "You do not have enough permissions";
    }
}