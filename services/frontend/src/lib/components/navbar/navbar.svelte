<script lang="ts">
  import { _ } from "svelte-i18n";
</script>

<header>
  <div class="branding">
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
  </div>
  <nav>
    <ul>
      <li>
        <a href="/">{$_("page.navigation.home")}</a>
      </li>
      <li>
        <a
          href="https://github.com/uvulpos/golang-sveltekit-binary"
          target="_blank">{$_("page.navigation.github")}</a
        >
      </li>
    </ul>
  </nav>
</header>

<style lang="sass">
  header
    display: flex
    justify-content: space-between
    padding: .2rem 1rem
    background-color: #ccc

    .branding
      display: flex
      align-items: center
      gap: .5rem

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
          color: #000
          font-weight: bold


</style>
