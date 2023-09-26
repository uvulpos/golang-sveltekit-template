<script lang="ts">
  import { Alert } from "$lib/components/alert";
  import { setUserProfile } from "$lib/stores/userProfile";
  import axios from "axios";
  import { onMount } from "svelte";

  onMount(() => {
    handleLogout();
    setUserProfile(undefined);
  });

  async function handleLogout() {
    try {
      const loginRequest = await axios.post(
        "/api/logout",
        {},
        {
          withCredentials: true,
        }
      );

      const statusCode = loginRequest.status;
      console.log("Login:", { statusCode });
    } catch (error) {
      console.log("error: ", { error });
    }
  }
</script>

<Alert>
  Logout! <br />
  Bye 👋🏻
</Alert>
