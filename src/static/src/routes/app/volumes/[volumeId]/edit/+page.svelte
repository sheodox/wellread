<style>
	textarea {
		width: 100%;
		height: 30rem;
	}
</style>

<Breadcrumbs {links} />
<form method="POST" use:enhance action="?/update" class="f-column gap-4 mt-4">
	<label class="blocky-label"
		>Name
		<br />
		<input name="name" value={data.volume.name} />
	</label>
	<div class="f-row f-wrap gap-4 f-1">
		<label class="blocky-label">
			Current page
			<br />
			<input name="currentPage" bind:value={currentPage} on:input={onCurrentPageChange} type="number" />
		</label>
		<label class="blocky-label">
			Newly read pages
			<br />
			<input name="pagesRead" bind:value={pagesRead} type="number" />
		</label>
		<label class="blocky-label">
			Status
			<br />
			<div>
				<select name="status" bind:value={status}>
					{#each readingStatuses as s}
						<option value={s}>{s.at(0)?.toUpperCase() + s.substring(1)}</option>
					{/each}
				</select>
				<ReadingStatusBadge {status} />
			</div>
		</label>
	</div>
	<label class="blocky-label"
		>Notes
		<br />
		<textarea name="notes" value={data.volume.notes} />
	</label>
	<div class="f-row">
		<a class="button secondary" href="/app/volumes/{data.volume.id}">Cancel</a>
		<button class="primary">Save</button>
	</div>
</form>

<script lang="ts">
	import { enhance } from '$app/forms';
	import Breadcrumbs from '$lib/Breadcrumbs.svelte';
	import ReadingStatusBadge from '$lib/ReadingStatusBadge.svelte';
	import { readingStatuses } from '$lib/types';
	import type { PageData } from './$types';
	export let data: PageData;

	let currentPage = data.volume.currentPage,
		status = data.volume.status,
		pagesRead = 0;

	function onCurrentPageChange() {
		pagesRead = Math.max(0, currentPage - data.volume.currentPage);
	}

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
			href: `/app/volumes/${data.volume.id}`,
			text: data.volume.name,
		},
		{ text: 'Update' },
	];
</script>
