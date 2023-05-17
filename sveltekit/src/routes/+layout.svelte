<script>
  import Footer from "../lib/components/layout/footer.svelte";
  import Header from "../lib/components/layout/header.svelte";
</script>

<div class="root">
  <header>
    <Header />
  </header>
  <slot />
  <footer>
    <Footer />
  </footer>
</div>

<style lang="scss">
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    font-weight: 400;
  }
  :global(p, h1, h2, h3, h4, h5) {
    margin: 0;
  }
  .root {
    display: flex;
    flex-direction: column;
    min-height: 100vh;

    footer {
      margin-top: auto;
    }
  }
</style>
