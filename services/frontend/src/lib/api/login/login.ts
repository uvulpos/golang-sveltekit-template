
export async function loginUserByCredentials(username: string, password: string): Promise<string | null> {
    try {
        let request: Response
        let statusCode: number
        let responseText: string
        request = await fetch("/api/login", {
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
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
