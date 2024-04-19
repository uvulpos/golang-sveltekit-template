<script lang="ts">
  //   import { Icon } from "../icon";
  import Logo from "$lib/assets/logo.svg";
  import "@fontsource-variable/roboto-condensed";

  export let roundCorner: boolean = false;
</script>

<nav class:round-corner={roundCorner}>
  <ul class="content">
    <li>
      <div class="header">
        <div class="logo">
          <img src={Logo} alt="Logo" />
        </div>
        <div class="search">
          <form>
            <input placeholder="search..." />
          </form>
        </div>
      </div>
    </li>
    <li>
      <ul class="navigation">
        <li><a href="/">Home</a></li>
        <li><a href="/about">About</a></li>
      </ul>
    </li>
  </ul>
</nav>

<style lang="sass">
    @import "src/lib/styles/_main.scss"
    
    nav
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
                    img
                        height: 3rem
                .search
                    input
                        border: none
                        border-bottom: 1.5px solid $ui-navigationbar-bg-secondary
                        background-color: transparent
                        color: #ffffff
                        outline: 0
                        padding: .5rem .75rem
            ul.navigation
                list-style-type: none
                background-color: $ui-navigationbar-bg-secondary
                display: flex
                gap: 2rem
                color: $ui-navigationbar-text-secondary
                padding: .5rem 5%
                font-family: 'Roboto Condensed Variable',Arial,sans-serif
                font-weight: 400
                line-height: 1.5
                font-size: 1rem
                
                a
                    text-decoration: none
                    color: $ui-navigationbar-text-secondary

</style>
