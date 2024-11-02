<script lang="ts">
  import { page } from "$app/stores";
  import { logoutSession } from "$lib/functions/logout/logout";
  import type { Page } from "@sveltejs/kit";
  import { _ } from "svelte-i18n";
  import { slide } from "svelte/transition";

  let activeElementIndex = 0;

  const navigation = [
    {
      name: $_("page.navigation.home"),
      href: "/",
      onclickFn: undefined,
    },
    {
      name: "Settings",
      href: "/settings",
      onclickFn: undefined,
      subelements: [
        {
          name: "Swagger",
          href: "/swagger",
          onclickFn: undefined,
        },
      ],
    },
  ];

  $: changeNavigationActiveElement($page);

  function changeNavigationActiveElement(
    page: Page<Record<string, string>, string | null>
  ) {
    navigation.forEach((navElement, index) => {
      const webPath = page.url.pathname + "/";
      const webPathRegex = new RegExp(`${navElement.href}\/.*`);
      const match = webPathRegex.test(webPath);

      if (match) {
        activeElementIndex = index;
        return;
      }
    });
  }
</script>

<ul class="navigation">
  {#each navigation as nav, counter}
    {@const activeElement = counter === activeElementIndex}
    <li>
      <a class:active-element={activeElement} href={nav.href}>{nav.name}</a>
      {#if activeElement}
        <div class="subelements" transition:slide>
          {#if nav.subelements}
            {#each nav.subelements as sub}
              <div class="subelement">
                <a href={sub.href}>{sub.name}</a>
              </div>
            {/each}
          {/if}
        </div>
      {/if}
    </li>
  {/each}
  <li style="margin-top: auto;"></li>
</ul>

<style lang="sass">

    ul.navigation
        list-style-type: none
        display: flex
        flex-direction: column
        // gap: .5rem
        margin: 0
        padding: 0
        margin-top: 3vh

        li
            a 
                text-decoration: none
                color: var(--sidebar-font-color)
                padding: .5rem 1rem
                display: block
                &:hover
                  background-color: var(--sidebar-background-color-hover)
                &.active-element
                  background-color: var(--sidebar-background-color-active)
            
            .subelements
              display: flex
              flex-direction: column
              background-color: var(--sidebar-background-color-submenu)

              a
                padding-left: 1.5rem
                &:hover
                  background-color: var(--sidebar-background-color-submenu-hover)

</style>
