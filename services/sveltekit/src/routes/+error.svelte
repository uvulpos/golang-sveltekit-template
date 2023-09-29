<script>
    import { _ } from 'svelte-i18n'
  
  </script>

<div class="content">
    <h1>Error 404</h1>
    <p>{$_('error.http.404.text')}</p>
</div>

<style lang="scss">
    .content {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%,-50%);
    }
</style>