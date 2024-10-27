<script lang="ts">
  import { onMount } from "svelte";
  import { invalidateAll } from "$app/navigation";

  onMount(() => {
    const interval = setInterval(() => {
      invalidateAll();
      window.location.href = "/api/v1/auth/logout";
    }, 1000);

    return () => {
      clearInterval(interval);
    };
  });
</script>
