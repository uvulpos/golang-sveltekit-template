<script lang="ts">
  import { toast } from "$lib/stores/toast/index";
  import { onMount } from "svelte";
  import ToastItem from "./toastItem.svelte";
  import { backInOut } from "svelte/easing";
  import { fade } from "svelte/transition";

  //   onMount(() => {
  //     toast.push("hello world, how are you");
  //     toast.push("moin");
  //   });
</script>

<ul>
  {#each $toast as item (item.id)}
    <li
      in:fade={{ duration: 200, easing: backInOut }}
      out:fade={{ duration: 200, easing: backInOut }}
    >
      <ToastItem {item} />
    </li>
  {/each}
</ul>

<style lang="sass">
    ul
        display: flex
        flex-direction: column
        gap: 1rem
        list-style-type: none
        margin: 0
        padding: 0
        li
            width: 100%
</style>
