<script lang="ts">
  import { generateHash } from "$lib/functions/random/gen-hash";
  import { Button } from "@svelteuidev/core";
</script>

<h1>Login to Application</h1>
<div class="actions">
  <Button href={"/api/v1/oauth/redirect?hash=" + generateHash(8)}>Login</Button>
</div>

<style lang="sass">
  .actions
    margin-top: 1rem
</style>
