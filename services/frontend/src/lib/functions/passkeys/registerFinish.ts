import { loginUserByWebAuthNFinish } from "$lib/api/login/fido-register-finish";

interface PasskeyCredentials {
    id: string;
    type: string;
    response: PasskeyCredentialsResponse;
    rawId: ArrayBuffer;
}

interface PasskeyCredentialsResponse {
    attestationObject: ArrayBuffer;
    clientDataJSON: ArrayBuffer;
}

type credtialType = Credential | null | undefined
export async function loginWithWebAuthNFinish(originalCredential: credtialType) {

    if (originalCredential == null) {
        return;
    }

    const credentials = originalCredential as PasskeyCredentials;
    const credentialResponse = credentials.response;

    const response = {
        id: credentials.id,
        rawId: arrayBufferToUrlBase64(credentials.rawId),
        type: credentials.type,
        response: {
            attestationObject: arrayBufferToUrlBase64(credentialResponse.attestationObject),
            clientDataJSON: arrayBufferToUrlBase64(credentialResponse.clientDataJSON),
        },
    };

    const result = await loginUserByWebAuthNFinish(response);
    return result
}

function arrayBufferToUrlBase64(buffer: ArrayBuffer): string {
    let binary = "";
    const bytes = new Uint8Array(buffer);
    const len = bytes.byteLength;
    for (let i = 0; i < len; i++) {
        binary += String.fromCharCode(bytes[i]);
    }
    return btoa(binary)
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=/g, "");
}