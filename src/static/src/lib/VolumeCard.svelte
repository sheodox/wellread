<style lang="scss">
	.volume-info {
		color: var(--sx-gray-100);
	}
</style>

<LinkCard href="/app/volumes/{volume.id}">
	<div class="volume p-3">
		<p class="volume-title sx-font-size-6 m-0 fw-bold"><ReadingStatusBadge status={volume.status} /> {volume.name}</p>
		<p class="volume-info sx-font-size-4 m-0">
			Page {volume.currentPage}
			{#if showSeries} - {volume.seriesName}{/if}
		</p>
	</div>
</LinkCard>

<script lang="ts">
	import LinkCard from './LinkCard.svelte';
	import ReadingStatusBadge from './ReadingStatusBadge.svelte';
	import type { Volume } from '$lib/types';

	export let volume: Volume;
	export let showSeries = true;
</script>
