<script>
  import Errorpage from "$lib/components/errorpage/errorpage.svelte";
  import { jwtDataStore, jwtStore } from "$lib/stores/jwt/jwt";
  import { _ } from "svelte-i18n";
</script>

{#if $jwtStore != null}
  <div class="root">
    <div class="navigation">
      <div class="navigation-panel">
        <ul>
          <li>
            <a href="/dashboard">
              <img src="/assets/vector/home.svg" alt="Home house" />
              <span>{$_("page.navigation.dashbord")}</span>
            </a>
          </li>
          <li>
            <a href="/settings">
              <img src="/assets/vector/grinding-gear.svg" alt="Settings Gear" />
              <span>{$_("page.navigation.settings")}</span>
            </a>
          </li>
          {#if $jwtDataStore?.permissions.includes("GREET_ADMIN")}
            <li>
              <a href="/settings/admin-settings">
                <img src="/assets/vector/user-tie.svg" alt="Admin User" />
                <span>{$_("page.navigation.admin-settings")}</span>
              </a>
            </li>
          {/if}
          <li>
            <a href="/logout">
              <img src="/assets/vector/logout.svg" alt="Logout" />
              <span> {$_("page.navigation.logout")} </span>
            </a>
          </li>
        </ul>
      </div>
    </div>
    <div class="content">
      <slot />
    </div>
  </div>
{:else}
  <Errorpage errorcode={401} />
{/if}

<style lang="sass">
    @import "../../lib/variables/sass/main"
    .root
        padding: 0 2rem 0 0 
        display: flex
        gap: 2rem
        width: 100%
        .navigation
            background-color: $ui-background-secondary
            min-width: clamp(200px, 20ch, 25vw)
            .navigation-panel
                border-radius: 5px
                padding: 1rem .5rem
                display: flex
                flex-direction: column
                justify-content: space-between
                gap: .5rem
                ul
                    padding: 0
                    list-style-type: none
                    display: flex
                    flex-direction: column
                    gap: .2rem 0

                    a
                        text-decoration: none
                        color: $ui-font-color
                        font-weight: bolder
                        display: flex
                        align-items: center
                        gap: .5rem
                        padding: .5rem .5rem 
                        &:hover
                          background-color: darken($ui-background-secondary, 3%)
                        img 
                          filter: invert(1)
                          height: 1.3rem
                          width: 1.3rem
        
        .content
            padding: 2rem 0
            flex-grow: 1
            max-width: 1000px
            margin: 0 auto
              


</style>
