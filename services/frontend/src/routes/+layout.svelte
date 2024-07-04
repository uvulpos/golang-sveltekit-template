<script lang="ts">
  // svelte config
  export const prerender = false;
  export const ssr = false;
  export const trailingSlash = "always";

  // import code
  import { onMount } from "svelte";
  import { _ } from "svelte-i18n";
  import { addMessages, getLocaleFromNavigator, init } from "svelte-i18n";
  import {
    AppShell,
    Navbar,
    Header,
    SvelteUIProvider,
  } from "@svelteuidev/core";
  import { DarkSidebar, DarkNavbar } from "./style";

  // install fonts
  import "@fontsource/inter";

  // import i18n files
  import en from "$lib/i18n/en.json";
  import de from "$lib/i18n/de.json";

  // configure i18n
  addMessages("en", en);
  addMessages("de", de);
  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });

  let opened = true;
  let preMount: boolean = true;

  onMount(() => {
    preMount = false;
  });
</script>

<SvelteUIProvider withNormalizeCSS withGlobalStyles>
  {#if !preMount}
    <AppShell>
      <Navbar
        slot="navbar"
        hidden={!opened}
        fixed
        width={{ sm: 250, lg: 250 }}
        override={DarkSidebar}
      >
        <span>Header 1</span>
      </Navbar>
      <Header slot="header" fixed height={60} override={DarkNavbar}>
        <span>Header 2</span>
      </Header>
      <slot />
    </AppShell>
  {/if}
</SvelteUIProvider>
