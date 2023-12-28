<script lang="ts">
  import { checkEmail, checkUsername } from "$lib/api/self";
  import { changeUserPassword } from "$lib/api/login/change-password";
  import { updateUser } from "$lib/api/self";
  import { getUser, type SelfUser } from "$lib/api/self/getUser";
  import Banner from "$lib/components/banner/banner.svelte";
  import Button from "$lib/components/button/button.svelte";
  import { CopyInput } from "$lib/components/input";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { jwtDataStore } from "$lib/stores/jwt/jwt";
  import { toast } from "$lib/stores/toast";
  import { onMount } from "svelte";
  import { _ } from "svelte-i18n";

  // requested data
  let selfUserAccount: Promise<SelfUser | null>;

  // inputs
  let oldPassword: string = "";
  let password: string = "";
  let passwordRepeat: string = "";

  onMount(() => {
    selfUserAccount = getUser();
  });

  // async function changeUserData() {
  //   selfUserAccount = updateUser();
  // }

  async function changePassword() {
    if (password !== passwordRepeat) {
      toast.push(
        "error",
        "Error",
        "could not change password. They are not equal"
      );
    }
    let username: string = "";
    if ($jwtDataStore !== undefined) {
      username = $jwtDataStore?.username;
    }

    await changeUserPassword(password, oldPassword);
    toast.push("success", "Success!", "Password got changed");
  }
</script>

<svelte:head>
  <title>{$_("page.titles.settings")}</title>
</svelte:head>

<div class="header">
  <div class="content">
    <p>Settings</p>
  </div>
</div>

<br />
<br />

{#await selfUserAccount}
  <p>Loading...</p>
{:then selfUser}
  <div>
    {#if selfUser?.auth_source !== "basic"}
      <Banner
        text="You are not using the integrated authentication tool. To update your
        personal information, you must use another tool / ask your administrator"
        type="warning"
      ></Banner>
    {/if}
  </div>
  <div class="input-form">
    <h2>General Information</h2>
    <form>
      <div class="flex-gap">
        <CopyInput value={selfUser?.uuid} showDefaultMargin={false} />
        <div class="grid-2-rows">
          <Textinput
            labelName="Username"
            value={selfUser?.username}
            disabled={selfUser?.auth_source !== "basic"}
            showDefaultMargin={false}
            validateFunction={checkUsername}
          />
          <Textinput
            labelName="Email"
            value={selfUser?.email}
            disabled={selfUser?.auth_source !== "basic"}
            showDefaultMargin={false}
            validateFunction={checkEmail}
          />
        </div>
        <div>
          <Button>Save</Button>
        </div>
      </div>
    </form>
  </div>

  <br />

  <div class="input-form">
    <h2>Change Password</h2>
    <form>
      <div class="flex-gap">
        <div>
          <Textinput
            disabled={selfUser?.auth_source !== "basic"}
            type="password"
            autocomplete="old-password"
            bind:value={oldPassword}
            labelName="Old Password"
            showDefaultMargin={false}
            placeholder="old password"
          />
        </div>
        <div class="grid-2-rows">
          <Textinput
            disabled={selfUser?.auth_source !== "basic"}
            type="password"
            autocomplete="new-password"
            bind:value={password}
            labelName="Password"
            showDefaultMargin={false}
            placeholder="password"
          />
          <Textinput
            disabled={selfUser?.auth_source !== "basic"}
            type="password"
            autocomplete="off"
            bind:value={passwordRepeat}
            labelName="Password again"
            showDefaultMargin={false}
            placeholder="repeat password"
          />
        </div>
        <div>
          <Button on:click={changePassword}>Save</Button>
        </div>
      </div>
    </form>
  </div>
  <br />
{:catch error}
  <p>Error: {error}</p>
{/await}

<!-- 
    if account is not of type basic dont make it deletable
 -->
<!-- <div>
  <h2>Delete Account</h2>
  <div>
    <Button>Delete Account</Button>
  </div>
</div> -->

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
      .input-form
        .flex-gap
          display: flex
          flex-direction: column
          gap: 1.5rem 
          .grid-2-rows
            display: grid
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr))
            gap: 1.5rem
</style>
