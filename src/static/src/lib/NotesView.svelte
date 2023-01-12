<style>
	article {
		border-radius: 3px;
	}
	.no-notes {
		font-style: italic;
	}
	.notes {
		white-space: pre-line;
	}
</style>

<article class="py-1">
	<h2 class="blocky-label">Notes</h2>
	{#if notes}
		<p class="notes m-0">{notes}</p>
	{:else}
		<p class="no-notes m-0">You haven't written any notes yet.</p>
	{/if}
</article>

<script lang="ts">
	export let notes: string;
</script>
