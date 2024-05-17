<script lang="ts">
  //   import { Icon } from "../icon";
  import Logo from "$lib/assets/logo.svg";
  import "@fontsource-variable/roboto-condensed";
  import { Icon } from "../icon";

  export let enableSearch: boolean = false;
  export let mobileVersion: boolean = false;
  export let environment: string = "";
  export let roundCorner: boolean = false;

  let showMobileNavigation: boolean = false;
</script>

<nav class:round-corner={roundCorner}>
  <ul class="content">
    <li>
      <div class="header">
        <div class="logo">
          <img src={Logo} alt="Logo" />
          {#if environment !== ""}
            <div class="environment">
              <p>{environment}</p>
            </div>
          {/if}
        </div>
        {#if enableSearch}
          <div class="search">
            <form>
              <input placeholder="search..." />
              <Icon icon="magnify-glass"></Icon>
            </form>
          </div>
        {/if}
      </div>
    </li>
    <li>
      <div class="navigation">
        <div class="icon" class:showHamburgerMenu={mobileVersion}>
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <!-- svelte-ignore a11y-missing-attribute -->
          <!-- svelte-ignore a11y-no-static-element-interactions -->
          <a
            on:click={() => {
              showMobileNavigation = !showMobileNavigation;
            }}
          >
            <Icon icon="hamburger-menu"></Icon>
          </a>
        </div>
        <ul
          class="navigation-elements"
          class:mobile-navigation-elements={mobileVersion}
          class:showMobileNavigation
        >
          <li><a href="/">Home</a></li>
          <li><a href="/about">About</a></li>
        </ul>
      </div>
    </li>
  </ul>
</nav>

<style lang="sass">
    @import "../../styles/_main.scss"

    * 
      margin: 0
      padding: 0
    
    nav
        font-family: sans-serif
        &.round-corner
            border-radius: 10px
            overflow: hidden
        ul.content
            list-style-type: none
            display: flex
            flex-direction: column
            padding: 0
            margin: 0
            .header
                padding: 1rem 5%
                background-color: $ui-navigationbar-bg-primary
                color: #ffffff
                display: flex
                justify-content: space-between
                align-items: center
                .logo
                    display: flex
                    flex-direction: row
                    align-items: center
                    gap: 2rem
                    img
                      height: 3rem
                    .environment
                      p
                        margin: 0
                        padding: 0
                        color: red
                        font-size: 1.5rem
                        font-weight: bolder
                .search
                    background-color: lighten($ui-navigationbar-bg-primary, 10%)
                    border-radius: 10px
                    padding: .2rem
                    form
                      transition: all 200ms ease-in-out
                      input
                          border: none
                          border-bottom: 1.5px solid $ui-navigationbar-bg-secondary
                          background-color: transparent
                          color: #ffffff
                          outline: 0
                          padding: .5rem .75rem


            .navigation
              background-color: $ui-navigationbar-bg-secondary
              color: $ui-navigationbar-text-secondary
              padding: .5rem 5%
              font-weight: 400
              line-height: 1.5
              font-size: 1rem

              .icon
                display: none
                width: 100%
                text-align: center
                &.showHamburgerMenu
                  display: block


              ul.navigation-elements
                  list-style-type: none
                  display: flex
                  gap: 2rem
                  padding: 0
                  transition: max-height 500ms ease-in-out
                  a
                      text-decoration: none
                      color: $ui-navigationbar-text-secondary

                  &.mobile-navigation-elements
                    flex-direction: column
                    gap: 0rem
                    max-height: 0

                    li 
                      a
                        display: block
                        font-size: 1.3rem
                        padding: .5rem 0
                        width: 100%

                    &.showMobileNavigation
                      max-height: 100vh
</style>
