<script>
  import { SvelteUIProvider } from "@svelteuidev/core";
  import "$lib/theme/index.sass";
</script>

<SvelteUIProvider withNormalizeCSS withGlobalStyles>
  <div class="background">
    <div class="content">
      <slot />
    </div>
  </div>
</SvelteUIProvider>

<style lang="sass">
    :global(body)
        margin: 0
        padding: 0
    .background
        width: 100%
        height: 100%
        display: block
        position: fixed
        background-image: linear-gradient( 73.1deg,  #5f7eb0 8%, #22334f 86.9%) !important
    .content
        position: absolute
        top: 40%
        transform: translateY(-50%)
        background: var(--background-color)
        display: inline-block
        margin-left: 15vw
        padding: 2rem
        border-radius: 10px
</style>
