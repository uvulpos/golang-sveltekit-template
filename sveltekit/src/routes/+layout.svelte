<script>
  export const prerender = false;
  export const ssr = false;
  export const trailingSlash = "always";

  import Footer from "../lib/components/layout/footer.svelte";
  import Header from "../lib/components/layout/header.svelte";

  import { addMessages, getLocaleFromNavigator, init } from "svelte-i18n";

  import en from "./../lib/i18n/en.json";
  import de from "./../lib/i18n/de.json";

  addMessages("en", en);
  addMessages("de", de);

  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });
</script>

<div class="root">
  <header>
    <Header />
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
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    font-weight: 400;
    font-size: 16px;
    background-color: rgb(18, 18, 22);
    color: #ffffff9b;
  }
  :global(p, h1, h2, h3, h4, h5) {
    margin: 0;
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
