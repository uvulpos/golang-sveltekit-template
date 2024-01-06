
export type RespondWithStatus = {
    "status": boolean;
}

export async function checkEmail(email: string): Promise<boolean | null> {
    try {
        let request: Response
        let statusCode: number
        request = await fetch("/api/v1/login/is-available-email", {
            method: "POST",
            body: JSON.stringify({
                email: email,
            }),
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status
        const responseJson = await (await awaitedResponse.json() as Promise<RespondWithStatus>)

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return false
        }

        return responseJson.status

    } catch (error) {
        return null
    }
}