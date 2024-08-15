type SelfPermission = {
    username: string;
    profilePicture: string;
    permissions: string[];
};

export async function checkUserPermission(fetch: (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>, jwt: string, requiredPermissions: string[]): Promise<boolean> {
    let hasPermission = true;
    try {
        const request = fetch(`/api/v1/self/permissions`, {
            method: 'GET',
            headers: {
                'Cookie': `jwt=${jwt}`
            }
        });
        const response = await request;

        if (!response.ok) {
            return false;
        }

        const responseData: SelfPermission = await response.json();

        requiredPermissions.forEach(reqPermission => {
            console.log(reqPermission, responseData.permissions);

            if (!responseData.permissions.includes(reqPermission)) {
                hasPermission = false;
            }
        });

        return hasPermission;
    } catch (error) {
        // Handle network errors or other unexpected errors
        return false;
    }
}