<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let message = 'Hello World';

	const dispatch = createEventDispatcher();

	function handleClick() {
		dispatch('click');
	}
</script>

<div class="example">
	<p>{message}</p>
	<button on:click={handleClick}>Click me</button>
</div>

<style>
	.example {
		padding: 1rem;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	button {
		margin-top: 0.5rem;
		padding: 0.5rem 1rem;
		background-color: #007bff;
		color: white;
		border: none;
		border-radius: 4px;
		cursor: pointer;
	}

	button:hover {
		background-color: #0056b3;
	}
</style>
