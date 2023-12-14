<script lang="ts">
  import type { HTMLInputTypeAttribute } from "svelte/elements";

  export let name: string;
  export let autocomplete: string;
  export let type: HTMLInputTypeAttribute | null | undefined = "text";
  export let placeholder: string = "";
  export let value: string;
</script>

{#if type === "password"}
  <input {name} {autocomplete} type="password" {placeholder} bind:value />
{:else}
  <input {name} {autocomplete} type="text" {placeholder} bind:value />
{/if}

<style lang="sass">
    @import "../../variables/sass/main"

    input
        border: none
        padding: .5rem 1rem
        border-radius: 5px
        outline: none
        font-size: 1rem
        background-color: $ui-inputs-textinputs-background
        color: $ui-font-color
</style>
