{#if data.reading.data.length}
	<Fieldset legend="In Progress">
		<div class="f-row f-wrap">
			{#each data.reading.data as volume}
				<VolumeCard {volume} />
			{/each}
		</div>
	</Fieldset>
{/if}

<h2>Welcome to WellRead!</h2>
<p>
	WellRead can be your bookmark and your notebook when reading. Create a series and volumes from the menu to the left!
</p>

<p>WellRead is a work in progress by <a href="https://sheodox.com/" class="inline-link">sheodox</a>.</p>

<script lang="ts">
	import type { PageData } from './$types';
	import { Fieldset } from 'sheodox-ui';
	import VolumeCard from '$lib/VolumeCard.svelte';

	export let data: PageData;
</script>
