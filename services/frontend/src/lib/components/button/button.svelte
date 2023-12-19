<script lang="ts">
  type ButtonType = "button" | "submit" | "reset" | null | undefined;
  export let type: ButtonType = "button";
</script>

<button on:click {type}><slot /></button>

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

</style>
