
export async function refreshUserLoginJWT(username: string): Promise<string | null> {
    try {
        let request: Response
        let statusCode: number
        let responseText: string
        request = await fetch("/api/v1/refresh-jwt", {
            method: "POST",
            body: JSON.stringify({
                username: username,
            }),
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status
        responseText = await awaitedResponse.text()

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return null
        }

        return responseText

    } catch (error) {
        return null
    }
}
