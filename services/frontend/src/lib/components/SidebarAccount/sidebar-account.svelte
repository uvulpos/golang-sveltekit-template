<script lang="ts">
  import { Menu } from "@svelteuidev/core";
  import { Exit, Pencil1 } from "radix-icons-svelte";
  import { _ } from "svelte-i18n";
</script>

<div class="content-element sidebar-account-content-element">
  <div class="account">
    <img src="" alt="" class="profilepicture" />
    <span>uVulpos</span>
    <div class="margin-left">
      <Menu class="menu">
        <Menu.Label>{$_("page.navigation.application.header")}</Menu.Label>
        <a href="/settings">
          <Menu.Item icon={Pencil1} href="/settings"
            >{$_("page.navigation.application.settings")}</Menu.Item
          >
        </a>
        <Menu.Label>Account</Menu.Label>
        <a href="/api/v1/oauth/logout">
          <!-- <Menu.Item icon={Pencil1} href="/to/identity provider"
            >{$_("page.navigation.account.edit-account")}</Menu.Item
          > -->
          <Menu.Item icon={Exit} color="red"
            >{$_("page.navigation.account.logout")}</Menu.Item
          >
        </a>
      </Menu>
    </div>
  </div>
</div>

<style lang="sass">
  .content-element
    .account
      // border: 1px solid #000
      border-radius: 5px
      display: flex
      align-items: center
      gap: .5rem
      padding: .25rem .5rem

      .profilepicture
        $size: 30px
        border-radius: 100%
        display: block
        height: $size
        width: $size
        background-color: #f60

      .margin-left
        margin-left: auto

      :global(.menu button:hover)
        background-color: var(--sidebar-background-color-hover)
      :global(.menu button svg)
        color: var(--sidebar-svg-color)

</style>
