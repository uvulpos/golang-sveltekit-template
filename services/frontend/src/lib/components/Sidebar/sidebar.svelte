<script>
  import { _ } from "svelte-i18n";
  import { SidebarAccount } from "../SidebarAccount";
  import { SidebarNavigation } from "../SidebarNavigation";
  import { Collapsible } from "svelte-collapsible";
</script>

<ul class="sidebar-content">
  <li>
    <SidebarNavigation />
  </li>
  <li class="margin-top" />
  <li>
    <Collapsible>
      <div slot="header" class="collapsible-header">
        <span>Example Nav Item</span>
      </div>
      <ul class="subnavigation">
        <li>
          <a href="/subitem1">Subitem 1</a>
        </li>
        <li>
          <a href="/subitem2">Subitem 2</a>
        </li>
        <li>
          <a href="/subitem3">Subitem 3</a>
        </li>
        <li>
          <a href="/subitem4">Subitem 4</a>
        </li>
        <li>
          <a href="/subitem5">Subitem 5</a>
        </li>
      </ul>
    </Collapsible>
  </li>
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

  .collapsible-header
    background-color: var(--sidebar-collapsible-background-color)
    color: var(--sidebar-collapsible-font-color)
    padding: .5rem
    cursor: pointer

  ul.subnavigation
    list-style-type: none
    padding: 0
    margin: 0
    &>li
      a
        text-decoration: none
        color: var(--sidebar-collapsible-font-color)
        padding: .5rem
        display: block
        &:hover
          background-color: var(--sidebar-collapsible-background-color-hover)
</style>
