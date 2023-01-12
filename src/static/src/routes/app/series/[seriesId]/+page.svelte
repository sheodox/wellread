<style>
	button {
		margin: 0;
	}
	.series-data {
		display: flex;
	}
	@media (max-width: 800px) {
		.series-data {
			flex-direction: column;
		}
	}
</style>

<Breadcrumbs {links} />

<div class="series-data gap-6">
	<div class="f-4">
		<div class="f-row justify-content-between align-items-center">
			<h1>{data.series.name}</h1>
			<MenuButton>
				<span slot="trigger"><Icon icon="ellipsis-v" variant="icon-only" /><span class="sr-only">Menu</span></span>
				<ul slot="menu">
					<li><button on:click={confirmDelete}><Icon icon="trash" />Delete</button></li>
				</ul>
			</MenuButton>
		</div>
		<div class="f-row justify-content-end">
			<a href="/app/series/{data.series.id}/edit" class="button primary"><Icon icon="edit" /> Update</a>
		</div>
		<NotesView notes={data.series.notes} />
	</div>

	<div class="f-1">
		<h2 class="blocky-label mb-0">Volumes</h2>
		{#if !showNewVolume}
			<div class="f-column">
				<button class="secondary my-2" on:click={() => (showNewVolume = true)}>New Volume</button>
				{#each data.volumes as volume}
					<VolumeCard {volume} showSeries={false} />
				{/each}
			</div>
		{:else}
			<form method="POST" action="?/newVolume" use:enhance class="f-column gap-2">
				<label class="blocky-label"
					>Name
					<br />
					<!-- svelte-ignore a11y-autofocus -->
					<input name="name" autofocus />
				</label>
				<div class="f-row gap-2">
					<button type="button" on:click={() => (showNewVolume = false)} class="secondary"> Cancel </button>
					<button class="primary"> Save </button>
				</div>
			</form>
		{/if}
	</div>
</div>

<script lang="ts">
	import { enhance } from '$app/forms';
	import { Icon, MenuButton } from 'sheodox-ui';
	import Breadcrumbs from '$lib/Breadcrumbs.svelte';
	import NotesView from '$lib/NotesView.svelte';
	import VolumeCard from '$lib/VolumeCard.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';

	export let data: PageData;

	let showNewVolume = false;

	const links = [
		{
			href: '/app/series',
			text: 'Series',
		},
		{
			text: data.series.name,
		},
	];

	async function confirmDelete() {
		// todo, attempt to replace with a createConfirmModal
		if (
			confirm(
				`Are you sure you want to delete "${data.series.name}"? This will delete ${data.series.volumeCount} volumes and cannot be undone.`
			)
		) {
			await fetch(`/app/series/${data.series.id}`, { method: 'DELETE' });
			goto(`/app/series/`);
		}
	}
</script>
