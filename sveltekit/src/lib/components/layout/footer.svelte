<script>
  import { _ } from "svelte-i18n";
</script>

<div class="content">
  <div class="links">
    <a href="https://github.com/uvulpos/golang-sveltekit-binary" target="_blank"
      >👉 {$_("header.navigation.github")}</a
    >
  </div>
  <div class="creator">
    <span>2023 @uvulpos</span>
  </div>
</div>

<style lang="scss">
  .content {
    margin-top: 3rem;
    background-color: rgb(0, 0, 0);
    color: white;
    display: block;
    padding: 2rem 2rem;

    .links {
      display: flex;
      justify-content: center;
      gap: 0.5rem;
      margin-bottom: 0.5rem;

      a {
        text-decoration: none;
        color: #fff;
      }
    }

    .creator {
      text-align: center;
      color: #7e7e7e;
    }
  }
</style>
