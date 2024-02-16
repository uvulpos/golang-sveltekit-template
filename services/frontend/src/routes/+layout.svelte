<script lang="ts">
  // import code
  import { onMount } from "svelte";
  import { _ } from "svelte-i18n";
  import { loginUserFromCookie } from "$lib/stores/jwt/jwt";
  import { addMessages, getLocaleFromNavigator, init } from "svelte-i18n";
  import "@fontsource/inter";

  // import components
  import { Navbar } from "$lib/components/navbar";
  import { Footer } from "$lib/components/footer";

  // import i18n
  import en from "$lib/i18n/en.json";
  import de from "$lib/i18n/de.json";

  // svelte config
  export const prerender = false;
  export const ssr = false;
  export const trailingSlash = "always";

  // configure i18n
  addMessages("en", en);
  addMessages("de", de);
  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });

  let preMount: boolean = true;

  onMount(() => {
    loginUserFromCookie();
    preMount = false;
  });
</script>

{#if !preMount}
  <div class="root">
    <header>
      <Navbar />
    </header>
    <div class="content">
      <!-- <div class="toasts">
        <Toast />
      </div> -->
      <slot />
    </div>
    <footer>
      <Footer />
    </footer>
  </div>
{/if}

<style lang="sass">
  @import "../lib/variables/sass/main"

  :global(body) 
    margin: 0
    padding: 0
    font-size: 16px
    background-color: $ui-background
    color: $ui-font-color
    font-family: "Inter"
  
  .root 
    display: flex
    flex-direction: column
    min-height: 100vh

    .content
      flex-grow: 1
      z-index: 8500
      display: flex
      justify-content: center
      .toasts
        position: fixed 
        z-index: 1000 
        right: 0
        margin-right: 3vw

    header
      box-shadow: 3px 3px 8px rgba(0,0,0,.2)
      z-index: 9000
    footer 
      margin-top: auto
</style>
