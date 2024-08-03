<div class="layout">
  <slot />
</div>

<style lang="sass">
    .layout
        margin: auto
        padding: max(2rem)
        max-width: 1000px

        :global(a) 
          text-decoration: none
          color: var(--anchor-font-color)

</style>
