<h1>Volumes</h1>
{#each data.volumes as volume}
	<VolumeCard {volume} />
{/each}

<script lang="ts">
	import type { PageData } from './$types';
	import VolumeCard from '$lib/VolumeCard.svelte';

	export let data: PageData;
</script>
