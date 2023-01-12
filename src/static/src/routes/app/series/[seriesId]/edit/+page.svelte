<style>
	input,
	textarea {
		width: 100%;
	}

	textarea {
		height: 30rem;
	}
</style>

<Breadcrumbs {links} />
<form method="POST" use:enhance action="?/update" class="f-column gap-4 mt-4">
	<label class="blocky-label"
		>Name
		<br />
		<input name="name" value={data.series.name} />
	</label>
	<label class="blocky-label"
		>Notes
		<br />
		<textarea name="notes" value={data.series.notes} />
	</label>
	<div class="f-row">
		<a class="button secondary" href="/app/series/{data.series.id}">Cancel</a>
		<button class="primary">Save</button>
	</div>
</form>

<script lang="ts">
	import { enhance } from '$app/forms';
	import Breadcrumbs from '$lib/Breadcrumbs.svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	const links = [
		{
			href: '/app/series',
			text: 'Series',
		},
		{
			href: `/app/series/${data.series.id}`,
			text: data.series.name,
		},
		{ text: 'Update' },
	];
</script>
