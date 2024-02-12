export interface PublicKey {
    publicKey: PublicKeyCredentials
}

export interface PublicKeyCredentials {
    rp: PublicKeyCredentialsRP;
    user: PublicKeyCredentialsUser,
    challenge: string;
    pubKeyCredParams: PublicKeyCredentialParameters[];
    timeout: number;
    authenticatorSelection: AuthenticatorSelectionCriteria;
}

export interface PublicKeyCredentialsRP {
    name: string;
    id: string;
}

export interface PublicKeyCredentialsUser {
    name: string;
    displayName: string;
    id: string;
}

export async function loginUserByWebAuthNBegin(): Promise<PublicKey | null> {
    try {
        let request: Response
        let statusCode: number
        let responseText: PublicKey
        request = await fetch("/api/v1/u2f/register", {
            headers: new Headers({ 'content-type': 'application/json' }),
        })
        const awaitedResponse = await request
        statusCode = awaitedResponse.status
        responseText = await awaitedResponse.json() as PublicKey

        // handle errors
        if (!request.ok || statusCode < 200 || statusCode > 299) {
            return null
        }

        return responseText

    } catch (error) {
        return null
    }
}