<style>
	.toolbar {
		background: var(--sx-gray-transparent-dark);
		border-radius: 10px;
	}
</style>

<div class="toolbar f-row p-1 gap-1">
	<form class="f-row f-1 justify-content-between" method="GET">
		<div class="f-row gap-1">
			<TextInput bind:value={name} name="name">Name</TextInput>
			<div class="sx-toggles align-items-center">
				<p class="m-0">Reading Status:</p>
				{#each statuses as status}
					{@const id = `reading-status-filter-${status}`}
					<input bind:group={readingStatus} type="radio" {id} value={status} name="status" />
					<label for={id} title={status}>
						{#if status}
							<ReadingStatusBadge {status} />
						{:else}
							Any
						{/if}
					</label>
				{/each}
			</div>
		</div>
		<button class="primary"
			><Icon icon="magnifying-glass" variant="icon-only" /><span class="sr-only">Search</span></button
		>
	</form>
</div>

<script lang="ts">
	import { TextInput, Icon } from 'sheodox-ui';
	import ReadingStatusBadge from '$lib/ReadingStatusBadge.svelte';
	import { readingStatuses } from '$lib/types';

	export let name: string;
	export let readingStatus: string;

	const statuses = [...readingStatuses, ''];
</script>
