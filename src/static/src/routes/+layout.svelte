<!-- svelte-ignore css-unused-selector -->
<style global lang="scss">
	@import '../../../../node_modules/@fortawesome/fontawesome-free/css/all.min.css';

	.blocky-label {
		color: var(--sx-gray-100);
		text-transform: uppercase;
		font-size: var(--sx-font-size-3);
		margin: 0;
		font-weight: bold;
	}
	:root {
		--sx-hue-gray: 195 !important;
	}
</style>

<SheodoxUIStyles />

<slot />

<script lang="ts">
	import { SheodoxUIStyles } from 'sheodox-ui';
</script>
