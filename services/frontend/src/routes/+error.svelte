<script lang="ts">
  import { onMount } from "svelte";
  import Cookies from "js-cookie";
  import { redirect } from "@sveltejs/kit";
  import { goto } from "$app/navigation";

  onMount(() => {
    const jwt = Cookies.get("jwt");
    if (!jwt) {
      window.location.href = "/login";
    }
  });

  function returnPreviousPage() {
    history.back();
  }
</script>

<div class="error">
  <h1>Error 404</h1>
  <span>Not Found</span>

  <div class="buttons">
    <button on:click={returnPreviousPage}>Go Back</button>
    <button on:click={async () => goto("/")}>Go Home</button>
  </div>
</div>

<style lang="sass">
    :global(body)
        margin: 0
        height: 100vh
        width: 100vw
        position: relative
        overflow: hidden   
    .error
        transform: translate(-50%, -50%)
        top: 40%
        left: 50%
        position: absolute
        display: flex
        flex-direction: column
        gap: 2rem
        text-align: center

        h1 
            font-size: 5rem
            margin: 0
            padding: 0
        span
            font-size: 2rem
        .buttons
          display: flex
          justify-content: center
          gap: 1rem
          button
            cursor: pointer
            flex-grow: 1
        

</style>
