<script lang="ts">
  import Alert from "$lib/components/alert/alert.svelte";
  import { Button } from "$lib/components/button";
  import { Card } from "$lib/components/card";
  import { Input } from "$lib/components/input";
  import { userProfileStore } from "$lib/stores/userProfile";
  import axios from "axios";

  let username: string;
  let password: string;

  async function submitLogin() {
    try {
      const loginRequest = await axios.post(
        "/api/login",
        {
          username: username,
          password: password,
        },
        {
          withCredentials: true,
        }
      );

      const statusCode = loginRequest.status;
      console.log("Login:", { statusCode });
    } catch (error) {
      console.log("test");
    }
  }
</script>

{#if $userProfileStore !== undefined}
  <Alert>
    <span>Already logged in</span>
  </Alert>
{:else}
  <div class="m-auto card">
    <Card cssClass="w-96">
      <form class="flex flex-col space-y-6" action="/">
        <h3 class="text-xl font-medium text-gray-900 dark:text-white">Login</h3>
        <div class="space-y-2">
          <Input label="Username / E-Mail" bind:value={username} />
        </div>
        <div class="space-y-2">
          <Input bind:value={password} label="Password" />
        </div>
        <Button on:click={submitLogin}>Login</Button>
        <div class="flex items-start gap-2 flex-col">
          <a
            href="/login/reset-password"
            class="text-sm text-primary-700 hover:underline dark:text-primary-500"
            >Lost password?</a
          >
          <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
            Not registered? <a
              href="/login/signup"
              class="text-primary-700 hover:underline dark:text-primary-500"
              >Create account</a
            >
          </div>
        </div>
      </form>
    </Card>
  </div>
{/if}

<style lang="scss">
  .card {
    :global(& > div) {
      max-width: 60vw;
    }
  }
</style>
