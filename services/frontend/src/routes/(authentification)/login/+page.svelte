<script lang="ts">
  import Button from "$lib/components/button/button.svelte";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { loginUser } from "$lib/stores/jwt/jwt-functions";
  import { _ } from "svelte-i18n";

  let username: string = "uvulpos";
  let password: string = "123";

  let loginFailed: boolean = false;

  async function tryLoginUser() {
    const result = await loginUser(username, password);
    if (result === false) {
      password = "";
    }
    loginFailed = !result;
  }
</script>

<div class="loginform">
  <div class="form">
    <p>Login Failed: {loginFailed ? "true" : "false"}</p>
    {#if loginFailed}
      <div class="errormessage">
        <p>{$_("page.login.wrong-credientials")}</p>
      </div>
    {/if}

    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>
      <span>Username / E-Mail</span>
      <Textinput
        name="username"
        autocomplete="username"
        placeholder="username"
        bind:value={username}
      />
    </label>
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>
      <span>Password</span>
      <Textinput
        name="password"
        autocomplete="off"
        type="password"
        placeholder="password"
        bind:value={password}
      />
    </label>
    <div class="submit">
      <Button on:click={tryLoginUser} type="submit">Login</Button>
    </div>
  </div>
</div>

<style lang="sass">
    .loginform
        transform: translate(-50%,-50%)
        top: 50%
        left: 50%
        position: absolute
        .form
            width: 400px
            display: flex
            flex-direction: column
            gap: 1rem

            .errormessage
              background-color: rgba(255,0,0,.2)
              padding: 1rem
              border-radius: 5px
              p
                padding: 0
                margin: 0

            label
                display: flex
                flex-direction: column
                gap: .5rem

                input
                    border: none
                    padding: .3rem .7rem
                    border: 1px solid #fff
                    border-radius: 5px
                    outline: none
            .submit
                margin-top: .5rem


</style>
