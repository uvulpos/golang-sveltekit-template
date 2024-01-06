<script lang="ts">
  import { _ } from "svelte-i18n";

  type BannerTypes = "error" | "information" | "warning";

  export let header: string | undefined = undefined;
  export let text: string;
  export let type: BannerTypes = "error";

  $: typeClass = `banner-${type}`;
  $: typeHeader = getTypeHeader(type);

  function getTypeHeader(type: BannerTypes) {
    switch (type) {
      case "information":
        return $_("banners.information");
      case "warning":
        return $_("banners.warning");
      default:
        return $_("banners.error");
    }
  }
</script>

<div class="errormessage {typeClass}">
  <span class="header">{header !== undefined ? header : typeHeader}</span>
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

        .header
          font-size: 1.5rem
          font-weight: bolder

        p
            padding: 0
            margin: 0
</style>
