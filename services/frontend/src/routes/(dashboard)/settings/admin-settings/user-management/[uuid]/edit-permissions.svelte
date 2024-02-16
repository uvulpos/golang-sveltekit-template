<script lang="ts">
  import type { ManagedUser } from "$lib/api/self/getUserByUUID";

  export let userAccount: ManagedUser | null;
</script>

<div>
  <ul>
    {#if userAccount !== undefined && userAccount?.permissions !== undefined}
      {#each userAccount?.permissions as permissions}
        <li>
          <span>{permissions.name}</span>
        </li>
      {/each}
    {/if}
  </ul>
</div>
