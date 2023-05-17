<script lang="ts">
    export let disabled: boolean
</script>

<button {disabled} on:click>
    <slot />
</button>

<style lang="scss">
    button {
        border: none;
        padding: .5rem .8rem;
        background-color: #ff3e00;
        color: white;   
        border-radius: 5px;
        font-size: 1.2rem;
        cursor: pointer;

        &:hover {
            background-color: #d63600;
        }

        &:disabled {
            background-color: #7b1f00;
        }
    }
</style>