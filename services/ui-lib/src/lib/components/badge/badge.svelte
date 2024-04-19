<script lang="ts">
  export let label: string;
</script>

<div class="badge">
  <span>{label}</span>
</div>

<style lang="sass">
    .badge
        background-color: #f60
        color: #fff
        padding: .3rem .75rem
        display: inline-block
        border-radius: 10px
</style>
