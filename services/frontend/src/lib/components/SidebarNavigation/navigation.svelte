<script>
  import { _ } from "svelte-i18n";
</script>

<ul class="navigation">
  <li>
    <a href="/">{$_("page.navigation.home")}</a>
  </li>
  <li>
    <a href="/swagger">Swagger</a>
  </li>
  <li>
    <a href="/api/v1/oauth/redirect">Oauth</a>
  </li>
  <li>
    <a href="/api/v1/oauth/logout">Logout</a>
  </li>
</ul>

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
                // background-color: #ccc
                background-color: var(--sidebar-background-color-hover)
                display: block
                border-radius: 10px
</style>
