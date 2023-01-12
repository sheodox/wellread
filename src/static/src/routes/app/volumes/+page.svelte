<h1>Volumes</h1>
{#each data.volumes as volume}
	<VolumeCard {volume} />
{:else}
	<p>
		You haven't added any volumes yet. Create a <a href="/app/series" class="inline-link">series</a>, then add a volume
		from the page for that series.
	</p>
{/each}

<script lang="ts">
	import type { PageData } from './$types';
	import VolumeCard from '$lib/VolumeCard.svelte';

	export let data: PageData;
</script>
