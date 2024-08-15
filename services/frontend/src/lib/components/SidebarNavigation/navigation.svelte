<script>
  import { _ } from "svelte-i18n";
  import { Accordion, AccordionItem } from "svelte-collapsible";
</script>

<Accordion>
  <li>
    <a href="/">{$_("page.navigation.home")}</a>
  </li>
  <AccordionItem key={1}>
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
  </AccordionItem>
</Accordion>

<style lang="sass">
    ul.navigation
        list-style-type: none
        display: flex
        flex-direction: column
        gap: .5rem
        margin: 0
        padding: 0
        margin-top: 5vh

        li
            a 
                text-decoration: none
                color: var(--sidebar-font-color)
                padding: .5rem .75rem
                background-color: #ccc
                display: block
                border-radius: 10px

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
