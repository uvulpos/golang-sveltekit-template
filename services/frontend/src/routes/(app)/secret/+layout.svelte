<script lang="ts">
  /** @type {import('./$types').PageData} */
  export let data;
</script>

<pre>
    {JSON.stringify(data, null, 2)}
</pre>

{#if data.status === 200}
  <slot></slot>
{:else}
  <h1>Not authorized / not permitted</h1>
  <p>
    You need to be logged in and have the required permissions to access this
    page.
  </p>
{/if}
