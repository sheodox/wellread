<style lang="scss">
	.link-card {
		border-radius: 5px;
		transition: border-color 0.2s;
		min-width: 10rem;

		&:hover {
			background: var(--sx-gray-transparent);
		}
	}
</style>

<a {href}>
	<div class="link-card">
		<slot />
	</div>
</a>

<script lang="ts">
	export let href: string;
</script>
