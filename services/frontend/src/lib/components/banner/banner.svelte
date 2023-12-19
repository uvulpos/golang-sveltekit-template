<script lang="ts">
  import { _ } from "svelte-i18n";

  type BannerTypes = "error" | "information" | "warning";

  export let text: string = "";
  export let type: BannerTypes = "error";

  $: typeClass = `banner-${type}`;
  $: typeHeader = getTypeHeader(type);

  function getTypeHeader(type: BannerTypes) {
    switch (type) {
      case "information":
        return "Information";
      case "warning":
        return "Warning";
      default:
        return "Error";
    }
  }
</script>

<div class="errormessage {typeClass}">
  <span>{typeHeader}</span>
  <p>{text}</p>
</div>

<style lang="sass">
    .errormessage
        padding: 1rem
        border-radius: 5px

        &.banner-error 
            background-color: rgba(255,0,0,.2)
        &.banner-information
            background-color: rgba(60,175,252,.2)
        &.banner-warning
            background-color: rgba(252,236,60,.2)

        p
            padding: 0
            margin: 0
</style>
