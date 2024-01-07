<script lang="ts">
  import { jwtDataStore } from "$lib/stores/jwt/jwt";
  import { _ } from "svelte-i18n";
</script>

<header>
  <a href="/" class="branding" title={$_("page.navigation.goto-landingpage")}>
    <img
      src="/assets/gifs/gopher-dance.gif"
      class="mr-3 h-6 sm:h-9"
      alt="Dancing Gopher as a Logo"
    />
    <span
      class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
    >
      Go Svelte!
    </span>
  </a>
  <nav>
    <ul>
      {#if $jwtDataStore !== undefined}
        <li>
          <a href="/dashboard">{$_("page.navigation.dashbord")}</a>
        </li>
        <li>
          <div class="dropdown-menu">
            <div class="dropdown-header">
              <span>{$jwtDataStore.username}</span>
              <img
                src="/assets/vector/arrow-dropdown.svg"
                alt="Dropdown Arrow"
              />
            </div>
            <div class="dropdown-wrapper">
              <div class="dropdown">
                <ul>
                  <li>
                    <a href="/dashboard">
                      <img src="/assets/vector/home.svg" alt="Home house" />
                      <span>{$_("page.navigation.dashbord")}</span>
                    </a>
                  </li>
                  <li><hr /></li>
                  <li>
                    <a href="/settings">
                      <img
                        src="/assets/vector/grinding-gear.svg"
                        alt="Settings Gear"
                      />
                      <span>
                        {$_("page.navigation.settings")}
                      </span>
                    </a>
                  </li>
                  {#if $jwtDataStore?.permissions.includes("GREET_ADMIN")}
                    <li>
                      <a href="/admin-settings">
                        <img
                          src="/assets/vector/user-tie.svg"
                          alt="Admin User"
                        />
                        <span>
                          {$_("page.navigation.admin-settings")}
                        </span>
                      </a>
                    </li>
                  {/if}
                  <li>
                    <a href="/logout">
                      <img src="/assets/vector/logout.svg" alt="Logout" />
                      <span>
                        {$_("page.navigation.logout")}
                      </span>
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </li>
      {:else}
        <li>
          <a href="/">{$_("page.navigation.home")}</a>
        </li>
        <li>
          <a href="/login">{$_("page.navigation.login")}</a>
        </li>
      {/if}
    </ul>
  </nav>
</header>

<style lang="sass">
  @import "../../variables/sass/main"

  header
    display: flex
    justify-content: space-between
    padding: .2rem 1rem
    background-color: #272b2f

    a.branding
      display: flex
      align-items: center
      gap: .5rem
      color: $ui-font-color
      text-decoration: none

      img
        height: 2.5rem
        width: 2.5rem
        object-fit: contain
    
    nav
      &>ul
        display: flex
        gap: .5rem
        list-style-type: none

        a
          text-decoration: none
          font-weight: bold
          color: $ui-font-color

        .dropdown-menu
          position: relative

          .dropdown-header
            display: flex
            align-items: center

            img
              height: 1rem
              filter: invert(1)

          &:hover > .dropdown-wrapper
            display: block

          .dropdown-wrapper
            z-index: 99999
            display: none
            position: absolute
            padding-top: 10px
            right: 0
            .dropdown
              background-color: #272b2f
              border: 1px solid $ui-font-color
              border-radius: 5px
              padding: .8rem
              min-width: 150px
              ul
                padding: 0
                list-style: none
                white-space: nowrap
                display: flex
                flex-direction: column
                gap: .5rem

                a
                  display: flex
                  align-items: center
                  gap: .5rem
                  img 
                    filter: invert(1)
                    height: 1rem
                    width: 1rem


</style>
