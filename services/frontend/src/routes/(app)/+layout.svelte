<script lang="ts">
  // svelte config
  // export const prerender = true;
  // export const ssr = true;
  // export const trailingSlash = "always";

  /** @type {import('./$types').PageData} */
  export let data;

  // import code
  import { onDestroy, onMount } from "svelte";
  import { _ } from "svelte-i18n";
  import { addMessages, init } from "svelte-i18n";
  import {
    AppShell,
    Navbar,
    Header,
    SvelteUIProvider,
  } from "@svelteuidev/core";
  import { allowedLanguages, getLocale } from "$lib/i18n";
  import { Header as PageHeader } from "$lib/components/Header";
  import { DarkNavbar } from "./style";
  import Cookies from "js-cookie";

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
  import { goto } from "$app/navigation";
  import { getSelfInformation } from "$lib/api/user/get-self_information";
  import { refreshJwtToken } from "$lib/api/authentication/refresh-jwt_token";
  import type { SelfInformation } from "$lib/api/user/models/SelfInformation";

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
  let pageIsLoading: boolean = true;
  let selfInformation: SelfInformation | null = null;
  let refreshIntervalID: number | undefined;

  onMount(async () => {
    preMount = false;
    bodyElement = document.body;
    themeStore.isThemeSetOrAutoDetect(window);

    // remove hash from url
    const refreshHahName = "refresh-hash";
    const url = new URL(window.location.href);

    if (url.searchParams.has(refreshHahName)) {
      url.searchParams.delete(refreshHahName);
      history.replaceState(null, "", url.toString());
    }

    // check if jwt exists
    let jwtToken = Cookies.get("jwt");

    if (jwtToken === undefined || jwtToken === "") {
      // check jwt refresh token
      const newToken = await refreshJwtToken();

      if (newToken === null || newToken === "") {
        window.location.href = "/login";
        return;
      }
    }

    let selfData = await getSelfInformation();

    if (selfData === undefined || selfData == null) {
      goto("/login");
    }

    refreshIntervalID = setInterval(
      async () => {
        const newToken = await refreshJwtToken();
        if (newToken === undefined || newToken === "") {
          clearInterval(refreshIntervalID);
          window.location.href = "/login";
          return;
        }
      },
      1000 * 60 * 5 // 5 minutes
    );

    selfInformation = selfData;
    pageIsLoading = false;
  });

  onDestroy(() => {
    clearInterval(refreshIntervalID);
  });

  $: {
    changePageTheme(bodyElement, $themeStore);
  }
</script>

{#if pageIsLoading}
  <div>
    <p>Loading...</p>
  </div>
{:else if !pageIsLoading && selfInformation != null}
  <SvelteUIProvider withNormalizeCSS withGlobalStyles>
    {#if !preMount}
      <AppShell class={"app-shell"}>
        <Navbar
          slot="navbar"
          fixed
          class="sidebar {!collapseSidebar
            ? 'sidebar-expandSidebar'
            : 'collapsedSidebar'}"
        >
          <Sidebar user={selfInformation} />
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
{:else}
  <div>
    <p>Unexpexted Error</p>
  </div>
{/if}

<style lang="sass">
  @import "$lib/theme/variables.scss"
  
  :global(body)
      background: var(--background-color)
      color: var(--font-color)

  :global(.app-shell)
    height: 100vh
    overflow-y: scroll

  :global(.sidebar)
      overflow-y: scroll
      padding: 1rem 0 2rem
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
