<script lang="ts">
  import { _ } from "svelte-i18n";

  type HttpStatusCodeTypes = 401 | 403 | 404;
  export let errorcode: HttpStatusCodeTypes = 404;
</script>

<div class="content">
  <h1>{$_(`error.http.statuscode.${errorcode}.header`)}</h1>
  <p>{$_(`error.http.statuscode.${errorcode}.text`)}</p>
</div>

<style lang="scss">
  .content {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>
