<style>
	button {
		margin: 0;
	}
</style>

<div class="f-row justify-content-between align-items-center mb-1">
	<h1 class="mb-0">Series</h1>
	<button on:click={() => (showNewSeries = !showNewSeries)} class="secondary align-self-end">New Series</button>
</div>
{#if showNewSeries}
	<form method="POST" action="?/newSeries" class="f-column gap-2 my-4 align-self-center" use:enhance>
		<label class="blocky-label">
			New Series Name
			<br />
			<input name="name" />
		</label>
		<div class="f-row gap-2 align-items-center">
			<button type="button" class="secondary" on:click={() => (showNewSeries = false)}>Cancel</button>
			<button class="primary">Add</button>
		</div>
	</form>
{:else}
	{#each data.series as series}
		<SeriesCard {series} />
	{/each}
{/if}

<script lang="ts">
	import { enhance } from '$app/forms';
	import SeriesCard from '$lib/SeriesCard.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
	let showNewSeries = false;
</script>
