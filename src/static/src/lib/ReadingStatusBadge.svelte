<style>
	span {
		width: var(--sx-spacing-3);
		height: var(--sx-spacing-3);
		display: inline-block;
		border-radius: 50%;
	}
	.status-planning {
		background-color: var(--sx-orange-500);
	}
	.status-reading {
		background-color: var(--sx-blue-500);
	}
	.status-completed {
		background-color: var(--sx-green-500);
	}
	.status-dropped {
		background-color: var(--sx-red-500);
	}
</style>

<span class="status-{status}" title={statusCapitalized} />

<script lang="ts">
	export let status: string;

	$: statusCapitalized = status.at(0)?.toUpperCase() + status.substring(1);
</script>
