import { loginUserByWebAuthNBegin } from "$lib/api/login/fido-login-begin";

export async function loginWithWebAuthNBegin() {
    if (!window.PublicKeyCredential) {
        return;
    }

    const publicKeys = await loginUserByWebAuthNBegin();
    if (publicKeys == null) {
        return;
    }

    const credentials = publicKeys.publicKey;

    const publicKeyCredentialCreationOptions: PublicKeyCredentialCreationOptions =
    {
        challenge: bufferChallengeDecode(credentials.challenge),
        rp: credentials.rp,
        user: {
            ...credentials.user,
            id: bufferUserIdBase64Decode(credentials.user.id),
        },
        pubKeyCredParams: credentials.pubKeyCredParams,
        authenticatorSelection: credentials.authenticatorSelection,
        timeout: credentials.timeout,
        attestation: "direct",
    };

    const credential = await navigator.credentials.create({
        publicKey: publicKeyCredentialCreationOptions,
    });

    return credential;
}

function bufferUserIdBase64Decode(value: string): Uint8Array {
    return Uint8Array.from(atob(value), (c) => c.charCodeAt(0));
}

function bufferChallengeDecode(value: string): Uint8Array {
    return Uint8Array.from(value, (c) => c.charCodeAt(0));
}