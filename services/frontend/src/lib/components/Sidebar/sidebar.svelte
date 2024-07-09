<script>
  import { _ } from "svelte-i18n";
  import { Button } from "@svelteuidev/core";
  import { SidebarAccount } from "../SidebarAccount";
  import { SidebarNavigation } from "../SidebarNavigation";
</script>

<ul class="sidebar-content">
  <li>
    <SidebarNavigation />
  </li>
  <li class="margin-top" />
  <!-- <li class="margin-top no-anchor-underline">
    <a href="/swagger" target="_blank">
      <Button size="xs" fullSize>Swagger</Button>
    </a>
  </li> -->
  <li>
    <SidebarAccount />
  </li>
  <li>
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
      .legal
        display: flex
        align-items: center
        justify-content: space-around

        a 
          text-decoration: none
          color: var(--sidebar-font-color)
          padding: .5rem
          transition: 200ms opacity ease-in-out
          &:hover
            opacity: .7
</style>
