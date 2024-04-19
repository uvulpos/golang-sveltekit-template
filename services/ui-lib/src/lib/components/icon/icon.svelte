<script lang="ts">
  import type { IconType } from ".";

  export let icon: IconType | string;
  export let fallbackicon: string = "unknown-icon";
  $: iconPath = getIconPath(icon);

  function getIconPath(icon: string): string {
    return `/assets/icons/${icon}.svg`;
  }

  function handleError() {
    iconPath = getIconPath(fallbackicon);
  }
</script>

<img src={iconPath} on:error={handleError} alt="Icon" />

<style lang="sass">
  img 
    $iconSize: var(--size, 24px)
    width: $iconSize
    height: $iconSize

</style>
