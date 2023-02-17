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
	div :global(fieldset) {
		margin: 0;
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
		<Fieldset legend="Volumes">
			{#if !showNewVolume}
				<div class="f-column">
					<div class="f-row align-items-center">
						<button class="secondary my-2" on:click={() => (showNewVolume = true)}>New Volume</button>

						<MenuButton>
							<span slot="trigger"><Icon icon="sort" variant="icon-only" /><span class="sr-only">Sort Order</span></span
							>
							<ul slot="menu">
								{#each sortOptions as opt}
									{@const selected = sortType === opt.sort}
									<li>
										<button aria-pressed={selected} on:click={() => (sortType = opt.sort)} disabled={opt.disabled}>
											{#if selected}
												<Icon icon="check" />
											{/if}
											{opt.text}
										</button>
									</li>
								{/each}
							</ul>
						</MenuButton>
					</div>
					{#each sortedVolumes as volume}
						<VolumeCard {volume} showSeries={false} />
					{/each}
					<Pagination
						variant="minimal"
						page={data.volumes.page.pageNumber}
						as="a"
						makeHref={makeVolumePageHref}
						max={Math.ceil(data.volumes.page.totalItems / data.volumes.page.pageSize)}
					/>
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
		</Fieldset>
	</div>
</div>

<script lang="ts">
	import { enhance } from '$app/forms';
	import { Icon, MenuButton, Fieldset, Pagination } from 'sheodox-ui';
	import Breadcrumbs from '$lib/Breadcrumbs.svelte';
	import NotesView from '$lib/NotesView.svelte';
	import VolumeCard from '$lib/VolumeCard.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import type { Volume } from '$lib/types';

	export let data: PageData;

	const allVolumesNumeric = data.volumes.data.every(({ name }) => {
		return /^[\d.]+$/.test(name);
	});

	type SortType = `${'alpha' | 'num'}-${'asc' | 'desc'}`;
	let sortType: SortType = allVolumesNumeric ? 'num-asc' : 'alpha-asc';

	$: sortedVolumes = [...data.volumes.data].sort(sortVolumes(sortType));

	const sortOptions: { sort: SortType; text: string; disabled: boolean }[] = [
		{
			sort: 'alpha-asc',
			text: 'Alpha A-Z',
			disabled: false,
		},
		{
			sort: 'alpha-desc',
			text: 'Alpha Z-A',
			disabled: false,
		},
		{
			sort: 'num-asc',
			text: 'Num 1-9',
			disabled: !allVolumesNumeric,
		},
		{
			sort: 'num-desc',
			text: 'Num 9-1',
			disabled: !allVolumesNumeric,
		},
	];

	function sortVolumes(sort: typeof sortType) {
		return (a: Volume, b: Volume) => {
			switch (sort) {
				case 'alpha-asc':
					return a.name.localeCompare(b.name);
				case 'alpha-desc':
					return b.name.localeCompare(a.name);
				case 'num-asc':
					return parseInt(a.name, 10) - parseInt(b.name, 10);
				case 'num-desc':
					return parseInt(b.name, 10) - parseInt(a.name, 10);
			}
		};
	}

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

	function makeVolumePageHref(pageNum: number) {
		return `/app/series/${data.series.id}?page=${pageNum}`;
	}

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
