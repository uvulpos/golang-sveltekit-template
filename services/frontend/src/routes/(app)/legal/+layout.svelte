<div class="layout">
  <slot />
</div>

<style lang="sass">
    .layout
        margin: auto
        padding: max(2rem)
        max-width: 1000px

        :global(p)
          line-height: 1.5

        :global(a) 
          text-decoration: none
          color: var(--anchor-font-color)

</style>
