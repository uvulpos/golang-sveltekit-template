<script lang="ts">
  import { changeUserPassword } from "$lib/api/login";
  import type { SelfUser } from "$lib/api/self/getUser";
  import { Button } from "$lib/components/button";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { jwtDataStore } from "$lib/stores/jwt/jwt";
  import { toast } from "$lib/stores/toast";

  export let userAccount: SelfUser | null;
  let oldPassword: string = "";
  let password: string = "";
  let passwordRepeat: string = "";

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

<div class="input-form">
  <h2>Change Password</h2>
  <form>
    <div class="flex-gap">
      <div>
        <Textinput
          disabled={userAccount?.auth_source !== "basic"}
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
          disabled={userAccount?.auth_source !== "basic"}
          type="password"
          autocomplete="new-password"
          bind:value={password}
          labelName="Password"
          showDefaultMargin={false}
          placeholder="password"
        />
        <Textinput
          disabled={userAccount?.auth_source !== "basic"}
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

<style lang="sass">
  @import "../../../lib/variables/sass/main"
  
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
