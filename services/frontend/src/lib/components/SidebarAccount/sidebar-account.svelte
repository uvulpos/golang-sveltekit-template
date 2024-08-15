<script lang="ts">
  import { Menu } from "@svelteuidev/core";
  import { Exit, Pencil1, User } from "radix-icons-svelte";
  import { _ } from "svelte-i18n";
</script>

<div class="content-element">
  <div class="account">
    <img src="" alt="" class="profilepicture" />
    <span>uVulpos</span>
    <div class="margin-left">
      <Menu>
        <Menu.Label>{$_("page.navigation.application.header")}</Menu.Label>
        <a href="/settings">
          <Menu.Item icon={Pencil1} href="/settings"
            >{$_("page.navigation.application.settings")}</Menu.Item
          >
        </a>
        <Menu.Label>Account</Menu.Label>
        <a href="/logout">
          <Menu.Item icon={Exit} href="/logout" color="red"
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

</style>
