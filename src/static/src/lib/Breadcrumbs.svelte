<style>
	.divider {
		color: var(--sx-gray-75);
	}
</style>

<div class="breadcrumbs f-row f-wrap gap-1 align-items-center">
	{#each links as link, index}
		{#if index > 0}<span class="sx-font-size-1 divider"><Icon icon="angle-right" variant="icon-only" /></span>{/if}
		{#if link.href}
			<a href={link.href} class="inline-link">{link.text}</a>
		{:else}
			<span>{link.text}</span>
		{/if}
	{/each}
</div>

<script lang="ts">
	import { Icon } from 'sheodox-ui';
	export let links: { href?: string; text: string }[];
</script>
