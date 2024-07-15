<script lang="ts">
  import { Button } from "@svelteuidev/core";
</script>

<h1>Login to Application</h1>
<div class="actions">
  <Button href="/api/v1/oauth/redirect">Login</Button>
</div>

<style lang="sass">
  .actions
    margin-top: 1rem
</style>
