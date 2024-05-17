<script lang="ts">
  import "./button.scss";

  export let label: string = "";
  export let href: string | undefined = undefined;
  export let targetBlank: boolean = false;
</script>

<a {href} target={targetBlank ? "_blank" : "_self"} on:click>
  <button class="btn information-button">
    <div class="label">
      {label}
    </div>
  </button>
</a>
