<script lang="ts">
  // i18n
  import { _ } from "svelte-i18n";
  import { CopyInput } from "$lib/components/input";
  import { checkEmail, checkUsername, updateUser } from "$lib/api/self";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { refreshUserLogin } from "$lib/stores/jwt/jwt-functions";
  import { Button } from "$lib/components/button";
  import type { SelfUser } from "$lib/api/self/getUser";
  import { toast } from "$lib/stores/toast";
  import { createEventDispatcher } from "svelte";
  import type { ManagedUser } from "$lib/api/self/getUserByUUID";
  import { updateUserByUUID } from "$lib/api/self/updateUserByUUID";

  export let uuid: string | undefined = undefined;
  export let userAccount: SelfUser | ManagedUser | null;

  let userUuid = userAccount?.uuid ?? $_("error.general.unknown");
  let email = userAccount?.email ?? "";
  let username = userAccount?.username ?? "";
  let authSource = userAccount?.auth_source ?? $_("error.general.unknown");
  let roleName = userAccount?.role_name ?? $_("error.general.unknown");

  const dispatch = createEventDispatcher();

  function handleEmailUpdate(e: CustomEvent) {
    email = e.detail.value;
  }
  function handleUsernameUpdate(e: CustomEvent) {
    username = e.detail.value;
  }

  async function changeUserData(selfUser: SelfUser | null) {
    let updatedUsername: string;
    let updatedEmail: string;

    updatedUsername = username ?? selfUser?.username ?? "";
    updatedEmail = email ?? selfUser?.email ?? "";

    let selfUserAccount;
    if (uuid === undefined) {
      selfUserAccount = updateUser(updatedUsername, updatedEmail);
    } else {
      selfUserAccount = updateUserByUUID(uuid, updatedUsername, updatedEmail);
    }
    const newSelfUserAccount = await selfUserAccount;
    dispatch("updateSelfUserAccount", { selfUserAccount });

    if (newSelfUserAccount === null) {
      toast.push(
        "error",
        $_("banners.custom.not-update-userdata.header"),
        $_("banners.custom.not-update-userdata.body")
      );
      return;
    }

    const newSession = await refreshUserLogin(updatedUsername);
    if (!newSession) {
      toast.push(
        "error",
        $_("banners.custom.not-refresh-session.header"),
        $_("banners.custom.not-refresh-session.body")
      );
      return;
    }
    toast.push(
      "success",
      $_("banners.custom.successful-updated.header"),
      $_("banners.custom.successful-updated.body")
    );
  }
</script>

<div class="input-form">
  <h2>{$_("page.settings.general-information")}</h2>
  <form>
    <div class="flex-gap">
      <div class="grid-2-rows">
        <CopyInput
          value={userUuid}
          showDefaultMargin={false}
          labelName={$_("page.settings.uuid")}
        />
        <CopyInput
          value={roleName}
          showDefaultMargin={false}
          labelName={$_("page.settings.role-name")}
        />
      </div>
      <div class="grid-2-rows">
        <Textinput
          labelName={$_("page.settings.username")}
          placeholder={$_("page.settings.username")}
          value={username}
          disabled={authSource !== "basic"}
          showDefaultMargin={false}
          validateFunction={checkUsername}
          on:inputchange={handleUsernameUpdate}
        />
        <Textinput
          labelName={$_("page.settings.email")}
          placeholder={$_("page.settings.email")}
          value={email}
          disabled={authSource !== "basic"}
          showDefaultMargin={false}
          validateFunction={checkEmail}
          on:inputchange={handleEmailUpdate}
        />
      </div>
      <div>
        <Button on:click={async () => changeUserData(userAccount)}
          >{$_("page.settings.save")}</Button
        >
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
