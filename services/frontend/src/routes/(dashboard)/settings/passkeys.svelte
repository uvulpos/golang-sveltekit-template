<script lang="ts">
  import { Button } from "$lib/components/button";
  import { loginWithWebAuthNBegin } from "$lib/functions/passkeys/registerBegin";
  import { loginWithWebAuthNFinish } from "$lib/functions/passkeys/registerFinish";

  async function loginWithWebAuthN() {
    loginWithWebAuthNBegin()
      .then(loginWithWebAuthNFinish)
      .catch(() => {
        // todo implement notification
        console.log("failed!");
      });
  }
</script>

<Button on:click={loginWithWebAuthN}>Passkeys erstellen (hardcodet)</Button>
