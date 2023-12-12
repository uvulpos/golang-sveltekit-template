<script lang="ts">
  export let type: string;
</script>

<button on:click {type}><slot /></button>

<style lang="sass">
    @import "../../variables/sass/main"

    button
        border: none
        background-color: #f60
        border-radius: 5px
        padding: .5rem 1rem
        color: $ui-font-color
        font-size: 1rem
</style>
