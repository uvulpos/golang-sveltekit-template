<script lang="ts">
  import { _ } from "svelte-i18n";
  import { SidebarAccount } from "../SidebarAccount";
  import { SidebarNavigation } from "../SidebarNavigation";
  import type { SelfInformation } from "$lib/api/user/models/SelfInformation";

  export let user: SelfInformation;
</script>

<ul class="sidebar-content">
  <li>
    <SidebarNavigation />
  </li>
  <li class="margin-top" />
  <li class="account-section">
    <SidebarAccount {user} />
  </li>
  <li class="legal-section">
    <hr />
    <div class="legal">
      <a href="/legal/imprint">{$_("page.navigation.imprint")}</a>
      <a href="/legal/gdpr-notice">{$_("page.navigation.gdpr-notice")}</a>
    </div>
  </li>
</ul>

<style lang="sass">
  ul.sidebar-content
    list-style-type: none
    display: flex
    flex-direction: column
    margin: 0
    padding: 0
    height: 100%
    &>li
      &.margin-top
        margin-top: auto
      &.legal-section, &.account-section
        padding: 0 1rem
      .legal
        display: flex
        align-items: center
        justify-content: space-around
        user-select: none

        a 
          text-decoration: none
          color: var(--sidebar-font-color)
          padding: .5rem
          transition: 200ms opacity ease-in-out
          &:hover
            opacity: .7
</style>
