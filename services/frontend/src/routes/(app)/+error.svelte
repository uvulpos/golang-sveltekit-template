<script>
  import { page } from "$app/stores";
</script>

<h1>{$page?.status} - {$page?.error?.message}</h1>
<!-- <h1>Error 404</h1> -->
