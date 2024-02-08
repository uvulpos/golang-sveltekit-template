<script lang="ts">
  // i18n
  import { _ } from "svelte-i18n";
  // svelte imports
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  // components
  import Banner from "$lib/components/banner/banner.svelte";
  import { getUserByUUID, type ManagedUser } from "$lib/api/self/getUserByUUID";
  import GeneralInformation from "../../../general-information.svelte";
  import ChangePassword from "../../../change-passwords.svelte";
  import EditPermissions from "./edit-permissions.svelte";

  // requested data
  let user: Promise<ManagedUser | null>;

  onMount(() => {
    const { uuid } = $page.params;
    user = getUserByUUID(uuid);
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

{#await user}
  <p>Loading...</p>
{:then userAccount}
  <br />
  <br />
  <div>
    {#if userAccount?.auth_source !== "basic"}
      <Banner
        text={$_("page.settings.ldap-user-no-edit-text")}
        header={$_("page.settings.ldap-user-no-edit-header")}
        type="warning"
      ></Banner>
    {/if}

    <GeneralInformation {userAccount} on:updateSelfUserAccount={() => {}} />
    <br />
    <EditPermissions {userAccount} />
    <br />
    <ChangePassword {userAccount} />
    <br />
  </div>
{:catch error}
  <p>Could not load data. Try again later</p>
{/await}

<style lang="sass">
        @import "../../../../../../lib/variables/sass/main"
    
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
</style>
