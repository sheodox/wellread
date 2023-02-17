<style>
	main {
		width: 80rem;
		max-width: 100vw;
		margin: 0 auto;
		display: flex;
		flex-direction: column;
	}
</style>

<svelte:head>
	<title>WellRead</title>
</svelte:head>

<Header appName="WellRead" bind:menuOpen showMenuTrigger={data.loggedIn} href="/app">
	<Logo slot="logo" />
	<nav slot="headerEnd">
		<ul>
			<li>
				<NavDropdown showOpenIcon={false}>
					<span slot="button">
						<Icon icon="user-circle" variant="icon-only" />
						<span class="sr-only">User menu</span>
					</span>
					<div slot="menu">
						<ul>
							<!-- svelte-ignore a11y-click-events-have-key-events -->
							<li on:click|stopPropagation>
								<form method="POST" action="/app?/logout">
									<button><Icon icon="sign-out-alt" />Log out</button>
								</form>
							</li>
						</ul>
					</div>
				</NavDropdown>
			</li>
		</ul>
	</nav></Header
>

<div class="f-row">
	<Sidebar bind:menuOpen docked>
		<div slot="header" class="f-row align-items-center">
			<Logo />
			<h1 class="ml-1">WellRead</h1>
		</div>
		<nav class="sx-sidebar-simple-links">
			<nav class="sx-sidebar-simple-links">
				{#each links as link}
					<a href={link.href} class:active={currentPathname === link.href}><Icon icon={link.icon} /> {link.text}</a>
				{/each}
			</nav>
		</nav>
	</Sidebar>
	<div class="f-1">
		<main class="p-2"><slot /></main>
	</div>
</div>

<script lang="ts">
	import { Header, Sidebar, NavDropdown, Icon } from 'sheodox-ui';
	import Logo from './Logo.svelte';
	import type { PageData } from './$types';
	import { afterNavigate } from '$app/navigation';

	export let data: PageData;

	let menuOpen = false;

	let currentPathname = '';

	afterNavigate(({ to }) => {
		currentPathname = to?.url.pathname || '';
	});

	const links = [
		{ href: '/app', text: 'Home', icon: 'home' },
		{ href: '/app/series', text: 'Series', icon: 'list' },
		{ href: '/app/volumes', text: 'Volumes', icon: 'book' },
	];
</script>
