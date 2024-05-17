<script>
  import { IconButton } from "../lib/components/icon-button";
</script>

<h1>Welcome to SvelteKit</h1>

<p>
  Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>

<IconButton
  icon="home"
  label="Hello World"
  targetBlank={true}
  href="https://google.com"
  on:click={() => {
    console.log("test");
  }}
/>
