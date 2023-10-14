<script>
  import { _ } from "svelte-i18n";
</script>

<div class="content">
  <div class="center">
    <p>{$_("page.home.hello-world")}</p>
  </div>
</div>

<style lang="sass">
  .content
    padding: 30vh 
    text-align: center
    font-size: 35px
    font-weight: bolder
</style>
