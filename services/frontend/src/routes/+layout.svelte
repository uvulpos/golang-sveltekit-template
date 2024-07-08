<script lang="ts">
  // svelte config
  // export const prerender = true;
  // export const ssr = true;
  // export const trailingSlash = "always";

  // import code
  import { onMount } from "svelte";
  import { _ } from "svelte-i18n";
  import { addMessages, init } from "svelte-i18n";
  import {
    AppShell,
    Navbar,
    Header,
    SvelteUIProvider,
  } from "@svelteuidev/core";
  import { getLocale } from "$lib/i18n";
  import { Header as PageHeader } from "$lib/components/Header";
  import { DarkNavbar } from "./style";

  // install fonts
  import "$lib/theme/import-me.scss";
  import "@fontsource/inter";

  // import i18n files
  import en from "$lib/i18n/en.json";
  import de from "$lib/i18n/de.json";
  import { logo } from "$lib/assets";
  import { Sidebar } from "$lib/components/Sidebar";
  import type { PageData } from "./$types";

  // configure i18n
  addMessages("en", en);
  addMessages("de", de);
  init({
    fallbackLocale: "en",
    initialLocale: getLocale(["en", "de"]),
  });

  export let data: PageData;

  let collapseSidebar = false;
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
        fixed
        class="sidebar {!collapseSidebar
          ? 'expandSidebar'
          : 'collapsedSidebar'}"
      >
        <Sidebar />
      </Navbar>
      <Header slot="header" fixed override={DarkNavbar} height={67}>
        <PageHeader {logo} bind:collapseSidebar />
      </Header>
      <div class="subpage-content" class:expandSidebar={collapseSidebar}>
        <slot />
      </div>
    </AppShell>
  {/if}
</SvelteUIProvider>
