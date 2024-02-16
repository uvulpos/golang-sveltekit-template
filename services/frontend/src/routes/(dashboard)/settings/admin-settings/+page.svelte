<script lang="ts">
  // i18n
  import { _ } from "svelte-i18n";
</script>

<svelte:head>
  <title>{$_("page.titles.settings")}</title>
</svelte:head>

<div class="header">
  <div class="content">
    <p>{$_("page.admin-settings.admin-settings")}</p>
  </div>
</div>

<br /><br />

<div class="actions">
  <a href="/settings/admin-settings/user-management" class="user-management">
    <img src="/assets/vector/user.svg" alt="user" />
    <div>
      <span>{$_("page.admin-settings.user-management-link")}</span>
    </div>
  </a>
  <a href="/settings/admin-settings/user-management" class="role-management">
    <img src="/assets/vector/people-group.svg" alt="people group" />
    <div>
      <span>{$_("page.admin-settings.role-management-link")}</span>
    </div>
  </a>
  <a
    href="/settings/admin-settings/user-management"
    class="permission-management"
  >
    <img src="/assets/vector/key.svg" alt="key" />
    <div>
      <span>{$_("page.admin-settings.permission-management-link")}</span>
    </div>
  </a>
</div>

<style lang="sass">
    @import "../../../../lib/variables/sass/main"

    .header
        height: 5rem
        width: 100%
        background-color: $ui-background-secondary
        position: relative
        
        .content
            position: absolute
            transform: translate(-50%,-50%)
            top: 50%
            left: 50%
            p
                font-weight: bolder
                font-size: 2rem
    
    .actions
      display: grid
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr))
      gap: 2rem

      & > a
        background-color: $ui-background-secondary
        text-align: center
        padding: 1.5rem
        border-radius: 5px
        text-decoration: none
        color: $ui-font-color
        transition: opacity 200ms ease-in-out

        &:hover
          opacity: .6

        img
          height: 50px
          width: 50px
          margin-bottom: 1rem

</style>
