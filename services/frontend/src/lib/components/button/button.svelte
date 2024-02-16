<script lang="ts">
  type ButtonType = "button" | "submit" | "reset" | null | undefined;
  type ButtonSize = "default" | "small";
  export let type: ButtonType = "button";
  export let size: ButtonSize = "default";
</script>

<button on:click {type} class={size}><slot /></button>

<style lang="sass">
    @import "../../variables/sass/main"

    button
        border: none
        background-color: $ui-color-primary
        border-radius: 5px
        padding: .5rem 1rem
        color: $ui-font-color
        font-size: 1rem
        cursor: pointer
        &:hover
          color: darken($ui-font-color, 10%)
          background-color: darken($ui-color-primary, 10%)

        &.small
          padding: .35rem .65rem
          font-size: .7rem


</style>
