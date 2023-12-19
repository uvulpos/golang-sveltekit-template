<script>
  import Errorpage from "$lib/components/errorpage/errorpage.svelte";
  import { jwtStore } from "$lib/stores/jwt/jwt";
  import { _ } from "svelte-i18n";
</script>

{#if $jwtStore != null}
  <div class="root">
    <div class="navigation">
      <div class="navigation-panel">
        <ul>
          <li><a href="/dashboard">Dashboard</a></li>
          <li>{$_("mock.something")}</li>
          <li>{$_("mock.something-else")}</li>
          <li><hr /></li>
          <li><a href="/settings">Settings</a></li>
          <li><a href="/logout">Logout</a></li>
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
        padding: 0 2rem
        display: flex
        gap: 2rem
    .navigation
        margin: 2rem 0
        .navigation-panel
            border-radius: 5px
            border: 1px solid $ui-font-color
            padding: 1rem 2rem
            ul
                padding: 0
                list-style-type: none
                display: flex
                flex-direction: column
                gap: .5rem

                a
                    text-decoration: none
                    color: $ui-font-color
                    font-weight: bolder
    
    .content
        flex-grow: 1
        padding: 2rem 0


</style>
