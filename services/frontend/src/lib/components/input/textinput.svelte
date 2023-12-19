<script lang="ts">
  import type { HTMLInputTypeAttribute } from "svelte/elements";

  export let labelName: string | undefined = undefined;
  export let showDefaultMargin: boolean = true;

  export let name: string = "";
  export let autocomplete: string | null | undefined = undefined;
  export let type: HTMLInputTypeAttribute | null | undefined = "text";
  export let placeholder: string = "";
  export let value: string = "";
  export let disabled: boolean = false;
</script>

<label class:showDefultMargin={showDefaultMargin}>
  {#if labelName !== undefined}
    <span>{labelName}</span>
  {/if}

  {#if type === "password"}
    <input
      {name}
      {autocomplete}
      type="password"
      {placeholder}
      bind:value
      {disabled}
    />
  {:else}
    <input
      {name}
      {autocomplete}
      type="text"
      {placeholder}
      bind:value
      {disabled}
    />
  {/if}
</label>

<style lang="sass">
    @import "../../variables/sass/main"

    label
      display: flex
      flex-direction: column
      gap: .5rem
      &.showDefultMargin
        margin: 1rem
      input
          border: none
          padding: .5rem 1rem
          border-radius: 5px
          outline: none
          font-size: 1rem
          background-color: $ui-inputs-textinputs-background
          color: $ui-font-color
          &:disabled
            cursor: not-allowed
            opacity: .3
</style>
