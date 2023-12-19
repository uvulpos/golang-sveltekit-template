<script lang="ts">
  export let labelName: string | undefined = undefined;
  export let showDefaultMargin: boolean = true;

  export let value: string = "";
</script>

<label class:showDefultMargin={showDefaultMargin}>
  {#if labelName !== undefined}
    <span>{labelName}</span>
  {/if}

  <input type="text" bind:value disabled />
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
