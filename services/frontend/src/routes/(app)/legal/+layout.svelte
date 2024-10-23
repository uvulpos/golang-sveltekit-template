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
        
        :global(h1)
          border-bottom: 1px solid var(--header-font-border-color)
          padding-bottom: 1rem
          margin-bottom: 2rem

</style>
