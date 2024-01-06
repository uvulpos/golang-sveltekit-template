<script lang="ts">
  import Banner from "$lib/components/banner/banner.svelte";
  import Button from "$lib/components/button/button.svelte";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { _ } from "svelte-i18n";

  let username: string = "uvulpos";
  let password: string = "123";

  let loginFailed: boolean = false;

  async function tryLoginUser() {
    // const result = await loginUser(username, password);
    // if (result === false) {
    //   password = "";
    // }
    // loginFailed = !result;
  }
</script>

<div class="loginform">
  <div class="form">
    {#if loginFailed}
      <!-- <Banner text={$_("page.login.wrong-credientials")}></Banner> -->
    {/if}

    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>
      <span>{$_("page.password-forgotten.username")}</span>
      <Textinput
        name="username"
        autocomplete="username"
        placeholder="username"
        showDefaultMargin={false}
        bind:value={username}
      />
    </label>
    <div class="last-row">
      <div class="submit">
        <Button on:click={tryLoginUser} type="submit"
          >{$_("page.password-forgotten.login")}</Button
        >
      </div>
      <div>
        <a href="/login">{$_("page.password-forgotten.rather-login")}?</a>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
      @import "../../../lib/variables/sass/main"
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
  
              .last-row
                display: flex
                align-items: center
                justify-content: space-between
                gap: 1rem
                .submit
                  margin-top: .5rem
                a 
                  color: $ui-font-color
  

</style>
