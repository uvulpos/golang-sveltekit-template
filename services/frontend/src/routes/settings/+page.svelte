<script lang="ts">
  import { allowedLanguages, getLocale } from "$lib/i18n";
  import { setLocale } from "$lib/i18n/locale";
  import { NativeSelect } from "@svelteuidev/core";
  import { _ } from "svelte-i18n";

  let currentLanguage = getLocale(allowedLanguages);

  function setLanguage(e: Event) {
    const target = event?.target as HTMLSelectElement;
    const language = target.value;
    setLocale(language, allowedLanguages);
    location.reload();
  }
</script>

<h1>{$_("page.settings.header")}</h1>

<NativeSelect
  data={[
    { label: "English", value: "en" },
    { label: "Deutsch", value: "de" },
  ]}
  placeholder={$_("page.settings.language.placeholder")}
  label={$_("page.settings.language.label")}
  description={$_("page.settings.language.description")}
  value={currentLanguage ?? undefined}
  required
  on:change={setLanguage}
/>
