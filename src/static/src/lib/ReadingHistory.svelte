<style>
	.pages-read {
		color: var(--sx-green-300);
	}
</style>

<aside>
	<h2 class="blocky-label">Reading History</h2>

	<div class="f-column gap-4 mt-2">
		{#each history as hist}
			<div>
				<form method="POST" action="?/deleteReadingHistory" use:enhance>
					<div class="f-row justify-content-between sx-font-size-5">
						<span>{hist.currentPage}</span>
						{#if hist.pagesRead}
							<span class="pages-read">+{hist.pagesRead}</span>
						{/if}
					</div>
					<div class="f-row justify-content-between align-items-center">
						<span>{new Date(hist.createdAt).toLocaleDateString()}</span>
						<button class="small" title="Delete this reading history"><Icon icon="times" variant="icon-only" /></button>
					</div>
					<input type="hidden" name="readingHistoryId" value={hist.id} />
				</form>
			</div>
		{/each}
	</div>
</aside>

<script lang="ts">
	import { Icon } from 'sheodox-ui';
	import { enhance } from '$app/forms';
	import type { ReadingHistory } from '$lib/types';
	export let history: ReadingHistory[];
</script>
