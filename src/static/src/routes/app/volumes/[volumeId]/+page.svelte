<style>
	.current-page {
		align-self: start;
		border-radius: 5px;
	}

	.volume-data {
		display: flex;
	}
	@media (max-width: 800px) {
		.volume-data {
			flex-direction: column;
		}
	}
</style>

<Breadcrumbs {links} />

<div class="f-column gap-6">
	<div class="f-row justify-content-between align-items-center">
		<h1 class="mb-0 f-row align-items-center gap-2">
			<ReadingStatusBadge status={data.volume.status} /> <span>{data.volume.name}</span>
		</h1>
		<MenuButton>
			<span slot="trigger"><Icon icon="ellipsis-v" variant="icon-only" /><span class="sr-only">Menu</span></span>
			<ul slot="menu">
				<li><button on:click={confirmDelete}><Icon icon="trash" />Delete</button></li>
			</ul>
		</MenuButton>
	</div>

	<div class="gap-4 volume-data">
		<div class="f-4 f-column gap-4">
			<div class="f-row justify-content-between">
				<div class="current-page f-row gap-4">
					<div class="f-column">
						<p class="blocky-label m-0 sx-font-size-3 fw-bold">Current page</p>
						<p class="m-0 sx-font-size-8">
							<strong>{data.volume.currentPage}</strong>
						</p>
					</div>
				</div>
				<a href="/app/volumes/{data.volume.id}/edit" class="button primary align-self-center"
					><Icon icon="edit" />Update</a
				>
			</div>
			<NotesView notes={data.volume.notes} />
		</div>
		{#if data.readingHistory.length}
			<div class="f-1">
				<ReadingHistory history={data.readingHistory} />
			</div>
		{/if}
	</div>
</div>

<script lang="ts">
	import { Icon, MenuButton } from 'sheodox-ui';
	import Breadcrumbs from '$lib/Breadcrumbs.svelte';
	import NotesView from '$lib/NotesView.svelte';
	import ReadingHistory from '$lib/ReadingHistory.svelte';
	import ReadingStatusBadge from '$lib/ReadingStatusBadge.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';

	export let data: PageData;

	const links = [
		{
			href: '/app/series',
			text: 'Series',
		},
		{
			href: `/app/series/${data.volume.seriesId}`,
			text: data.volume.seriesName,
		},
		{
			text: data.volume.name,
		},
	];

	async function confirmDelete() {
		// todo, attempt to replace with a createConfirmModal
		if (confirm(`Are you sure you want to delete "${data.volume.name}"?`)) {
			await fetch(`/app/volumes/${data.volume.id}`, { method: 'DELETE' });
			goto(`/app/series/${data.volume.seriesId}`);
		}
	}
</script>
