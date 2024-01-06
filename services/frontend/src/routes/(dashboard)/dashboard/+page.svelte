<script lang="ts">
  import { _ } from "svelte-i18n";
</script>

<svelte:head>
  <title>{$_("page.titles.dashboard")}</title>
</svelte:head>

<div class="header">
  <div class="content">
    <p>{$_("page.dashboard.dashboard")}</p>
  </div>
</div>

<style lang="sass">
    @import "../../../lib/variables/sass/main"

    .header
        height: 5rem
        width: 100%
        background-color: $ui-background-secondary
        position: relative
        
        .content
            position: absolute
            transform: translate(-50%,-50%)
            top: 50%
            left: 50%
            p
                font-weight: bolder
                font-size: 2rem
</style>
