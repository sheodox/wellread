<style>
	main {
		width: 80rem;
		max-width: 100vw;
		margin: 0 auto;
		display: flex;
		flex-direction: column;
	}
</style>

<Header appName="WellRead" slim={false}>
	<div slot="logo">
		<Logo />
	</div>
</Header>

<main class="p-2">
	<p class="sx-font-size-9 text-align-center">Read without losing the plot</p>

	<p class="text-align-center">
		Well Read helps you organize your notes about books you're reading, so you're never lost when starting a new volume.
	</p>

	<div class="text-align-center">
		<a href="/auth" data-sveltekit-reload class="button primary">Sign in / Register</a>
	</div>
</main>

<script lang="ts">
	import { Header } from 'sheodox-ui';
	import Logo from './app/Logo.svelte';
</script>
