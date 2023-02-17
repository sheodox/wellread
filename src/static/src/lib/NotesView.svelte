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
	<Fieldset legend="Notes">
		{#if notes}
			<p class="notes m-0">{notes}</p>
		{:else}
			<p class="no-notes m-0 muted">You haven't written any notes yet.</p>
		{/if}
	</Fieldset>
</article>

<script lang="ts">
	import { Fieldset } from 'sheodox-ui';
	export let notes: string;
</script>
