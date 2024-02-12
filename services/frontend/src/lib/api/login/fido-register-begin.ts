export async function registerUserByWebAuthNBegin(): Promise<PublicKeyCredentialCreationOptions | null> {
    try {
        let request: Response
        let statusCode: number
        let responseText: PublicKeyCredentialCreationOptions
        request = await fetch("/api/v1/u2f/register", {
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status
        responseText = await awaitedResponse.json() as PublicKeyCredentialCreationOptions

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return null
        }

        return responseText

    } catch (error) {
        return null
    }
}