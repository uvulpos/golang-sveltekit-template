export async function changeUserPassword(password: string, oldPassword: string): Promise<boolean | null> {
    try {
        let request: Response
        let statusCode: number
        request = await fetch("/api/v1/login/change-password", {
            method: "POST",
            body: JSON.stringify({
                newPassword: password,
                oldPassword: oldPassword
            }),
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return false
        }

        return true

    } catch (error) {
        return null
    }
}