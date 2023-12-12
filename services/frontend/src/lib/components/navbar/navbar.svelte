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
          <a href="/dashboard">{$_("page.navigation.home")}</a>
        </li>
        <li>
          <span>{$jwtDataStore.username}</span>
        </li>
        <li>
          <a href="/logout">{$_("page.navigation.logout")}</a>
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
      ul
        display: flex
        gap: .5rem
        list-style-type: none

        a
          text-decoration: none
          font-weight: bold
          color: $ui-font-color


</style>
