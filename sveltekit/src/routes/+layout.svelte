<script lang="ts">
  import "../app.postcss";
  export const prerender = false;
  export const ssr = false;
  export const trailingSlash = "always";

  import en from "$lib/i18n/en.json";
  import de from "$lib/i18n/de.json";

  import { Navbar } from "$lib/components/navbar";
  import { Footer } from "$lib/components/footer";

  import { addMessages, getLocaleFromNavigator, init } from "svelte-i18n";
  import { onMount } from "svelte";
  import { getThemePreference } from "$lib/function/setTheme";
  import {
    setUserProfile,
    type userProfileType,
  } from "$lib/stores/userProfile";
  import { getUserProfile } from "$lib/function/getProfile";

  addMessages("en", en);
  addMessages("de", de);

  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });

  onMount(() => {
    getThemePreference();
    setSession();
  });

  async function setSession() {
    let user: userProfileType | undefined;
    const userPromise = await getUserProfile();
    try {
      user = userPromise;
      console.log(userPromise);
    } catch (error) {
      user = undefined;
    }
    setUserProfile(user);
  }
</script>

<div class="root bg-white dark:bg-gray-900">
  <header>
    <Navbar />
  </header>
  <!-- <span style="color: white;">
    {JSON.stringify($userProfileStore)}
  </span> -->
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
