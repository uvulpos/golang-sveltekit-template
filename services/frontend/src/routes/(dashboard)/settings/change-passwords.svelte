<script lang="ts">
  import { _ } from "svelte-i18n";
  import { changeUserPassword } from "$lib/api/login";
  import type { SelfUser } from "$lib/api/self/getUser";
  import { Button } from "$lib/components/button";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { toast } from "$lib/stores/toast";
  import type { ManagedUser } from "$lib/api/self/getUserByUUID";
  import { changeUserPasswordByUUID } from "$lib/api/login/change-password-by-uuid";

  export let uuid: string | undefined = undefined;
  export let userAccount: SelfUser | ManagedUser | null;
  let oldPassword: string = "";
  let password: string = "";
  let passwordRepeat: string = "";

  async function changePassword() {
    if (password !== passwordRepeat) {
      toast.push(
        "error",
        "Error",
        "could not change password. They are not equal."
      );
    }

    let pwChanged;
    if (uuid === undefined) {
      pwChanged = await changeUserPassword(password, oldPassword);
    } else {
      pwChanged = await changeUserPasswordByUUID(uuid, password, oldPassword);
    }

    if (pwChanged) {
      toast.push("success", "Success!", "Password got changed");
      oldPassword = "";
      password = "";
      passwordRepeat = "";
    } else {
      toast.push("error", "Error!", "Password could not get changed.");
    }
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
          labelName={$_("page.settings.old-password")}
          showDefaultMargin={false}
          placeholder={$_("page.settings.old-password")}
        />
      </div>
      <div class="grid-2-rows">
        <Textinput
          disabled={userAccount?.auth_source !== "basic"}
          type="password"
          autocomplete="new-password"
          bind:value={password}
          labelName={$_("page.settings.password")}
          showDefaultMargin={false}
          placeholder={$_("page.settings.password")}
        />
        <Textinput
          disabled={userAccount?.auth_source !== "basic"}
          type="password"
          autocomplete="off"
          bind:value={passwordRepeat}
          labelName={$_("page.settings.repeat-password")}
          showDefaultMargin={false}
          placeholder={$_("page.settings.repeat-password")}
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
