<script lang="ts">
  import { onMount } from "svelte";
  import Cookies from "js-cookie";

  onMount(() => {
    const jwt = Cookies.get("");
    if (!jwt) {
      window.location.href = "/login";
    }
  });
</script>

<div class="error">
  <h1>Error 404</h1>
  <span>Not Found</span>
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
        

</style>
