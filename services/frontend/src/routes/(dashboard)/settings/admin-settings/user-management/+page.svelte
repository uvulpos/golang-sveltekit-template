<script lang="ts">
  import { loginUserByCredentials, type User } from "$lib/api/users/get-users";
  import Button from "$lib/components/button/button.svelte";
  import { onMount } from "svelte";

  // i18n
  import { _ } from "svelte-i18n";

  let UsersPromise: Promise<User[] | null>;

  onMount(() => {
    UsersPromise = loginUserByCredentials();
  });
</script>

<svelte:head>
  <title>{$_("page.titles.settings")}</title>
</svelte:head>

<div class="header">
  <div class="content">
    <p>{$_("page.admin-settings.user-management.user-management")}</p>
  </div>
</div>

<br /><br />

{#await UsersPromise}
  <span>Loading...</span>
{:then users}
  {#if users != null}
    <table>
      <tr>
        <th>ID</th>
        <th>Username</th>
        <th>Email</th>
        <th>Auth Src</th>
        <th>Suspendet</th>
        <th>Role Name</th>
        <th></th>
      </tr>
      {#each users as currentUser}
        <tr>
          <td title={currentUser.uuid}>{currentUser.uuid.slice(0, 8)}...</td>
          <td>{currentUser.username}</td>
          <td>{currentUser.email}</td>
          <td>{currentUser.auth_source}</td>
          <td>{currentUser.suspendet}</td>
          <td>{currentUser.role_name}</td>
          <td>
            <a href="user-management/{currentUser.uuid}">
              <Button size="small">Edit</Button>
            </a>
          </td>
        </tr>
      {/each}
    </table>
  {:else}
    <span>null</span>
  {/if}
{:catch error}
  <span>error</span>
{/await}

<style lang="sass">
      @import "../../../../../lib/variables/sass/main"
  
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

      table 
        width: 100%
        border-collapse: collapse
        tr
          th
            background-color: $ui-background-secondary
            padding: .75rem .5rem
            text-align: left
            font-weight: bold
            
          td
            padding: .5rem
            color: darken($ui-font-color, 20%)

          &:nth-child(odd)
            td
              background-color: $ui-background-secondary
</style>
