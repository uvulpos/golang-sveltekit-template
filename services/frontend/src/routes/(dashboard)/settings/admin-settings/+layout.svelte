<script>
  import Errorpage from "$lib/components/errorpage/errorpage.svelte";
  import { jwtDataStore, jwtStore } from "$lib/stores/jwt/jwt";
</script>

{#if $jwtDataStore != null && $jwtDataStore.permissions.includes("MANAGE_USERS")}
  <slot />
{:else}
  <Errorpage errorcode={403} />
{/if}
