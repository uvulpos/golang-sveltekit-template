<script>
  import Button from "$lib/components/button/Button.svelte";
  import { InputTypes } from "$lib/components/input";
  import Input from "$lib/components/input/input.svelte";
</script>

<div class="content">
  <form>
    <h1>Login</h1>
    <div>
      <Input label="Username" inputType={InputTypes.Text} />
    </div>
    <div>
      <Input label="Password" inputType={InputTypes.Password} />
    </div>
    <div class="submit">
      <Button>Login</Button>
    </div>
    <div class="forgot-passowrd">
      <a href="/login/password-reset">I forgot my password</a>
    </div>
  </form>
</div>

<style lang="scss">
  .content {
    form {
      width: fit-content;
      margin: 10vh auto;
      border-radius: 10px;
      padding: 3rem;
      display: flex;
      flex-direction: column;
      gap: 1rem;
      background-color: rgb(29, 29, 32);

      h1 {
        text-align: center;
        margin-bottom: 1rem;
      }

      div.submit {
        margin-top: 1rem;
        text-align: center;
      }

      .forgot-passowrd {
        text-align: center;
        a {
          text-decoration: none;
          color: inherit;
          opacity: 0.4;
        }
      }
    }
  }
</style>
