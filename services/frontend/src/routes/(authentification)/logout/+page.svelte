<script lang="ts">
  import { goto } from "$app/navigation";
  import { jwtSessionDestroy } from "$lib/stores/jwt/jwt";
  import { onMount } from "svelte";

  onMount(() => {
    jwtSessionDestroy();
    goto("/");
  });
</script>

<svelte:head>
  <title>Logout...</title>
</svelte:head>

<p>Logout...</p>
