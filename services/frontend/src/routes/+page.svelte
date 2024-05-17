<script>
  import { _ } from "svelte-i18n";
  // import { Author } from "";
</script>

<svelte:head>
  <title>{$_("page.titles.home")}</title>
</svelte:head>

<div class="content">
  <div class="center">
    <p>{$_("page.home.welcome-human")}</p>
    <!-- <Author></Author> -->
  </div>
</div>

<style lang="sass">
  @import "../lib/variables/sass/main"

  .content
    padding: 30vh 
    text-align: center
    font-size: 35px
    font-weight: bolder
</style>
