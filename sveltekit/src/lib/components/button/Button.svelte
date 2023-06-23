<script lang="ts">
  import { createEventDispatcher } from "svelte";

  export let disabled: boolean = false;

  const dispatchEvent = createEventDispatcher();
  function handleOnClick(e: Event) {
    if (!disabled) {
      dispatchEvent("onclick", { event: e });
    }
  }
</script>

<button {disabled} on:click={handleOnClick}>
  <slot />
</button>

<style lang="scss">
  button {
    border: var(--ui-button-border, none);
    padding: var(--ui-button-padding, 0.5rem 1rem);
    background-color: var(--ui-button-bgcolor, #f60);
    color: var(--ui-button-color, white);
    border-radius: var(--ui-button-borderradius, 5px);
    font-size: var(--ui-button-fontsize, 1rem);
    cursor: var(--ui-button-cursor, pointer);
    transition: background-color 200ms ease-in-out;

    &:hover {
      background-color: var(--ui-button-bgcolor-hover, rgb(189, 47, 0));
      color: var(--ui-button-color-hover, white);
    }

    &:disabled {
      background-color: var(--ui-button-bgcolor-disabled, rgb(176, 99, 44));
      color: var(--ui-button-color-disabled, rgba(255, 255, 255, 0.711));
      cursor: default !important;
    }
  }
</style>
