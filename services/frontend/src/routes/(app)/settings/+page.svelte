<script lang="ts">
  import { allowedLanguages, getLocale } from "$lib/i18n";
  import { setLocale } from "$lib/i18n/locale";
  import { ThemeEnum, themeStore } from "$lib/stores";
  import { NativeSelect, theme } from "@svelteuidev/core";
  import { _ } from "svelte-i18n";

  let currentLanguage = getLocale(allowedLanguages);
  let currentTheme = $themeStore;

  $: updateThemeFromStore($themeStore);
  function updateThemeFromStore(theme: any) {
    if (theme == ThemeEnum.Darkmode) currentTheme = "darkmode";
    else if (theme == ThemeEnum.Lightmode) currentTheme = "lightmode";
  }

  function setLanguage() {
    setLocale(currentLanguage ?? "", allowedLanguages);
    location.reload();
  }

  function setTheme() {
    if (currentTheme == "darkmode") {
      themeStore.toDarkMode();
    } else if (currentTheme == "lightmode") {
      themeStore.toLightMode();
    }
  }
</script>

<h1>{$_("page.settings.device_header")}</h1>
<span>{$_("page.settings.device_description")}:</span>

<div class="menu">
  <NativeSelect
    data={[
      { label: "English", value: "en" },
      { label: "Deutsch", value: "de" },
    ]}
    placeholder={$_("page.settings.language.placeholder")}
    label={$_("page.settings.language.label")}
    description={$_("page.settings.language.description")}
    bind:value={currentLanguage}
    required
    on:change={setLanguage}
  />
  <NativeSelect
    data={[
      { label: "Darkmode", value: "darkmode" },
      { label: "Lightmode", value: "lightmode" },
    ]}
    placeholder={$_("page.settings.theme.placeholder")}
    label={$_("page.settings.theme.label")}
    required
    bind:value={currentTheme}
    on:change={setTheme}
  />
</div>

<style lang="sass">
  span
    margin-top: .75rem
    display: block
  .menu
    margin-top: 2rem
    display: flex
    flex-direction: column
    gap: 1.5rem
</style>
