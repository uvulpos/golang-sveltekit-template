<script lang="ts">
  import { onMount } from "svelte";

  onMount(async () => {
    window.location.href = "/api/v1/oauth/logout";
  });
</script>
