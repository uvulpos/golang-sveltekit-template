<script lang="ts">
  import { Card } from "flowbite-svelte";

  export let cssClass: string;
</script>

<Card class={cssClass}><slot /></Card>
