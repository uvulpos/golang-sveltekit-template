<script lang="ts">
  import Button from "$lib/components/button/button.svelte";
  import Textinput from "$lib/components/input/textinput.svelte";
  import { loginUser } from "$lib/stores/jwt/jwt-functions";

  let username: string = "uvulpos";
  let password: string = "123";
</script>

<div class="loginform">
  <div class="form">
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
      <Button on:click={() => loginUser(username, password)} type="submit"
        >Login</Button
      >
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
