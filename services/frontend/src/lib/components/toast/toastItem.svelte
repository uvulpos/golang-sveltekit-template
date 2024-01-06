<script lang="ts">
  import { toast, type ToastInformation } from "$lib/stores/toast/toastStore";
  import { onMount } from "svelte";
  export let item: ToastInformation;

  export const destroyAfterSeconds: number = 10;
  let intervalID: number = 0;

  function destroyComponent() {
    toast.pop(item.id);
  }

  onMount(() => {
    intervalID = setTimeout(() => {
      destroyComponent();
    }, destroyAfterSeconds * 1000);
  });
</script>

<div class="toast toast-type-{item.type}">
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div class="closing-button" on:click={() => destroyComponent()}>
    <img src="/assets/vector/x.svg" alt="closing X symbol" />
  </div>
  <h6>{item.header}</h6>
  <p>{item.body}</p>
</div>

<style lang="sass">
    @import "../../variables/sass/main"
    .toast
        padding: 1rem
        border-radius: 5px
        border: 1px solid white
        background-color: $ui-background-secondary
        position: relative
        min-width: 20vw

        &.toast-type-success
            border-color: green
            h6
                color: green
        &.toast-type-error
            border-color: red
            h6
                color: red

        .closing-button
            position: absolute
            right: 1rem
            cursor: pointer
            img
                filter: invert(1)
                height: .8rem
                width: .8rem
            &:hover
                opacity: .7
        h6, p
            margin: 0
            padding: 0 
            font-size: 1rem

</style>
