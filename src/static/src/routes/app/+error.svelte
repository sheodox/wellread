<style>
	.error-message {
		color: var(--sx-red-300);
		font-family: monospace;
	}
</style>

<p>Sorry, something went wrong.</p>

<p class="error-message">Error: {$page.error?.message}</p>

<script lang="ts">
	import { page } from '$app/stores';
</script>
