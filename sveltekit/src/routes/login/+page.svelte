<script lang="ts">
  import Alert from "$lib/components/alert/alert.svelte";
  import { Button } from "$lib/components/button";
  import { Card } from "$lib/components/card";
  import Checkbox from "$lib/components/checkbox/checkbox.svelte";
  import { Input } from "$lib/components/input";
  import { userProfileStore } from "$lib/stores/userProfile";
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
          <Input label="Username / E-Mail" />
        </div>
        <div class="space-y-2">
          <Input label="Password" />
        </div>
        <Button>Login</Button>
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
