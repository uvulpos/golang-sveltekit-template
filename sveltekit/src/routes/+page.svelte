<script>
  import { _ } from "svelte-i18n";
</script>

<div class="content">
  <div class="welcome">
    <div class="gopher">
      <img
        alt="Svelte Society Gopher"
        src="/assets/images/svelte-society-gopher.png"
      />
    </div>

    <div class="header">
      <h1>{$_("page.home.header-1")}</h1>
    </div>
    <div class="text">
      <p>{$_("page.home.text")}</p>
    </div>
  </div>
</div>

<style lang="scss">
  .content {
    max-width: 1200px;
    margin: auto;

    .welcome {
      .gopher {
        margin-top: 2rem;
        text-align: center;
        img {
          height: 40vh;
        }
      }
      .header {
        margin-top: 1rem;
        text-align: center;
        color: #f60;
      }
      .text {
        text-align: justify;
        padding: 2rem 5rem;
      }
    }
  }
</style>
