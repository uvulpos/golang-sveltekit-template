<script lang="ts">
  import { _ } from "svelte-i18n";
  import { Burger } from "@svelteuidev/core";

  export let logo: string;
  export let collapseSidebar: boolean;
</script>

<div class="header">
  <a href="/" class="branding">
    <img src={logo} alt="logo" height="45" />
    <span>{$_("page.navigation.branding-name")}</span>
  </a>

  <div class="header-right" />
  <div class="burger-menu">
    <Burger
      opened={!collapseSidebar}
      color="white"
      on:click={() => (collapseSidebar = !collapseSidebar)}
    />
  </div>
</div>

<style lang="sass">

    .header 
        display: flex
        align-items: center
        gap: 1rem
        height: 100%
        padding: 0 1rem
        color: var(--lm-header-font-color)
        background-color: var(--lm-header-background-color)

        a 
            text-decoration: none


        .header-right 
            margin-left: auto
        
        a.branding
            text-decoration: none
            color: var(--lm-header-font-color)
            display: flex
            align-items: center
            gap: 1rem

        .burger-menu 
            display: block

        @media (min-width: 1024px) 
            .burger-menu 
                display: none
  

</style>
