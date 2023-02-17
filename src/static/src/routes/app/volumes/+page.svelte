<h1>Volumes</h1>
<VolumeFilter name={data.volumes.filter.name} readingStatus={data.volumes.filter.status} />
{#if data.volumes.data.length > 0}
	<div class="my-4">
		{#each data.volumes.data as volume}
			<VolumeCard {volume} />
		{/each}
	</div>
	<Pagination
		page={data.volumes.page.pageNumber}
		max={Math.ceil(data.volumes.page.totalItems / data.volumes.page.pageSize)}
		as="a"
		makeHref={makeVolumeHref}
	/>
{:else}
	<p>No Results</p>
	<p class="muted">
		Create a <a href="/app/series" class="inline-link">series</a>, then add a volume from the page for that series.
	</p>
{/if}

<script lang="ts">
	import type { PageData } from './$types';
	import VolumeCard from '$lib/VolumeCard.svelte';
	import VolumeFilter from './VolumeFilter.svelte';
	import { Pagination } from 'sheodox-ui';

	export let data: PageData;

	function makeVolumeHref(pageNum: number) {
		return `/app/volumes/?page=${pageNum}&name=${encodeURIComponent(data.volumes.filter.name)}&status=${
			data.volumes.filter.status
		}`;
	}
</script>
