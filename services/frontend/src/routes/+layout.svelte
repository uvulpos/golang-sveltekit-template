<script lang="ts">
  import en from "$lib/i18n/en.json";
  import de from "$lib/i18n/de.json";

  import { Navbar } from "$lib/components/navbar";
  import { Footer } from "$lib/components/footer";

  import { addMessages, getLocaleFromNavigator, init } from "svelte-i18n";

  export const prerender = false;
  export const ssr = false;
  export const trailingSlash = "always";

  addMessages("en", en);
  addMessages("de", de);

  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });
</script>

<div class="root">
  <header>
    <Navbar />
  </header>
  <slot />
  <footer>
    <Footer />
  </footer>
</div>

<style lang="scss">
  :global(body) {
    margin: 0;
    padding: 0;
    font-size: 16px;
  }
  .root {
    display: flex;
    flex-direction: column;
    min-height: 100vh;

    footer {
      margin-top: auto;
    }
  }
</style>
