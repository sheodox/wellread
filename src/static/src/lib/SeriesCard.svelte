<style lang="scss">
	.series-info {
		color: var(--sx-gray-100);
	}
</style>

<LinkCard href="/app/series/{series.id}">
	<div class="series p-3">
		<p class="series-title sx-font-size-6 m-0 fw-bold">{series.name}</p>
		<p class="series-info sx-font-size-4 m-0">
			{series.volumeCount} Volumes
		</p>
	</div>
</LinkCard>

<script lang="ts">
	import LinkCard from './LinkCard.svelte';
	import type { Series } from '$lib/types';

	export let series: Series;
</script>
