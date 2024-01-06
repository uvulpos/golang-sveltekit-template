<script lang="ts">
  // i18n
  import { _ } from "svelte-i18n";
  // svelte imports
  import { onMount } from "svelte";
  // apis
  import { getUser, type SelfUser } from "$lib/api/self/getUser";
  // components
  import Banner from "$lib/components/banner/banner.svelte";
  import GeneralInformation from "./general-information.svelte";
  import ChangePassword from "./change-passwords.svelte";

  // requested data
  let selfUserAccount: Promise<SelfUser | null>;

  onMount(() => {
    selfUserAccount = getUser();
  });

  function updateSelfUserAccount(e: CustomEvent) {
    selfUserAccount = e.detail.selfUserAccount;
  }
</script>

<svelte:head>
  <title>{$_("page.titles.settings")}</title>
</svelte:head>

<div class="header">
  <div class="content">
    <p>{$_("page.settings.settings")}</p>
  </div>
</div>

{#await selfUserAccount}
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
  </div>

  <GeneralInformation
    {userAccount}
    on:updateSelfUserAccount={updateSelfUserAccount}
  />
  <br />
  <ChangePassword {userAccount} />
  <br />
{:catch error}
  <p>Could not load data. Try again later</p>
{/await}

<style lang="sass">
      @import "../../../lib/variables/sass/main"
  
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
