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
    theme,
  } from "@svelteuidev/core";
  import { allowedLanguages, getLocale } from "$lib/i18n";
  import { Header as PageHeader } from "$lib/components/Header";
  import { DarkNavbar } from "./style";

  // install fonts
  import "$lib/theme/index.sass";
  import "@fontsource/inter";

  // import i18n files
  import en from "$lib/i18n/en.json";
  import de from "$lib/i18n/de.json";
  import { logo } from "$lib/assets";
  import { Sidebar } from "$lib/components/Sidebar";
  import { themeStore } from "$lib/stores";
  import { changePageTheme } from "$lib/functions/theme/theme";

  // configure i18n
  addMessages("en", en);
  addMessages("de", de);
  init({
    fallbackLocale: "en",
    initialLocale: getLocale(allowedLanguages),
  });

  let collapseSidebar = true;
  let preMount: boolean = true;
  let bodyElement: HTMLElement | undefined;

  onMount(() => {
    preMount = false;
    bodyElement = document.body;
    themeStore.isThemeSetOrAutoDetect(window);
  });

  $: {
    changePageTheme(bodyElement, $themeStore);
  }
</script>

<SvelteUIProvider withNormalizeCSS withGlobalStyles>
  {#if !preMount}
    <AppShell>
      <Navbar
        slot="navbar"
        fixed
        class="sidebar {!collapseSidebar
          ? 'sidebar-expandSidebar'
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

<style lang="sass">
  @import "$lib/theme/variables.scss"
  
  :global(body)
      background: var(--background-color)
      color: var(--font-color)

  :global(.sidebar)
      overflow-y: scroll
      padding: 1rem 1rem 2rem
      z-index: 100
      position: fixed
      inset: 0 
      width: 0px
      display: none
      height: calc(100vh - var(--topbar-height)) 
      margin-top: var(--topbar-height) 

      color: var(--sidebar-font-color)
      background-color: var(--sidebar-background-color)
      border: none
      box-sizing: border-box
      font-size: 0.9em

  :global(.sidebar-expandSidebar)
      display: flex 
      width: var(--sidebar-width)

  @media (min-width: 1024px)
    :global(.sidebar)
        display: block
        width: var(--sidebar-width) 
    :global(.sidebar-expandSidebar)
        display: flex
        width: var(--sidebar-width)
</style>
