<script lang="ts">
  import { loginUserByWebAuthNBegin } from "$lib/api/login/fido-login-begin";
  import { Button } from "$lib/components/button";

  function bufferDecode(value: string): Uint8Array {
    return Uint8Array.from(atob(value), (c) => c.charCodeAt(0));
  }

  async function loginWithWebAuthNBegin() {
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
        challenge: bufferDecode(credentials.challenge),
        rp: credentials.rp,
        user: {
          ...credentials.user,
          id: bufferDecode(credentials.user.id),
        },
        pubKeyCredParams: credentials.pubKeyCredParams,
        authenticatorSelection: credentials.authenticatorSelection,
        timeout: credentials.timeout,
        attestation: "direct",
      };

    const credential = await navigator.credentials.create({
      publicKey: publicKeyCredentialCreationOptions,
    });
    console.log(credential);
    return credential;
  }

  function loginWithWebAuthNFinish(originalCredential: Credential | null | undefined) {
    if (originalCredential == null) {
        return;
      }

      const credentials: PasskeyCredentials =
        originalCredential as PasskeyCredentials;

      let attestationObject = credentials.response.attestationObject;
      let clientDataJSON = credentials.response.clientDataJSON;
      let rawId = credentials.rawId;

      await loginUserByWebAuthNFinish()
  }

  async function loginWithWebAuthN() {
    loginWithWebAuthNBegin().then(loginWithWebAuthNFinish)
  }

  interface PasskeyCredentials {
    id: string;
    type: string;
    response: PasskeyCredentialsResponse;
    rawId: string;
  }

  interface PasskeyCredentialsResponse {
    attestationObject: string;
    clientDataJSON: string;
  }
</script>

<Button on:click={loginWithWebAuthN}>Passkeys erstellen (hardcodet)</Button>
