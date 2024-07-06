<script>
  import { Button, colorScheme } from "@svelteuidev/core";
  import { _ } from "svelte-i18n";
</script>

<svelte:head>
  <title>{$_("page.titles.home")}</title>
</svelte:head>

<div>
  <p>{$_("page.home.welcome-human")}</p>
  <a href="/swagger">
    <Button color="green">Swagger</Button>
  </a>
</div>
