<div class="f-column align-items-center justify-content-center f-1">
	<Logo />
	<h1>WellRead</h1>
	<div id="firebaseui-auth-container" />
</div>

<svelte:head>
	<link type="text/css" rel="stylesheet" href="https://www.gstatic.com/firebasejs/ui/6.0.1/firebase-ui-auth.css" />
	<script src="https://www.gstatic.com/firebasejs/8.10.1/firebase-app.js"></script>
	<script src="https://www.gstatic.com/firebasejs/8.10.1/firebase-auth.js"></script>
	<script src="https://www.gstatic.com/firebasejs/ui/6.0.1/firebase-ui-auth.js"></script>
</svelte:head>

<script context="module" lang="ts">
	export const prerender = false;
</script>

<script lang="ts">
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';
	import Logo from '../app/Logo.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
	/**
	This whole thing should be rewritten once firebaseui supports firebase@9
	**/

	onMount(() => {
		if (!browser) {
			return;
		}
		//eslint-disable-next-line @typescript-eslint/no-explicit-any
		const firebase = (window as any).firebase;
		firebase.initializeApp(data.firebaseConfig);

		// FirebaseUI config.
		const uiConfig = {
			signInSuccessUrl: '/api/auth/callback',
			callbacks: {
				signInSuccessWithAuthResult: function () {
					const user = firebase.auth().currentUser;

					if (user) {
						authUser(user);
					}

					return false;
				},
			},

			signInOptions: [
				// Leave the lines as is for the providers you want to offer your users.
				firebase.auth.GoogleAuthProvider.PROVIDER_ID,
				//firebase.auth.FacebookAuthProvider.PROVIDER_ID,
				//firebase.auth.TwitterAuthProvider.PROVIDER_ID,
				//firebase.auth.GithubAuthProvider.PROVIDER_ID,
				firebase.auth.EmailAuthProvider.PROVIDER_ID,
				//firebase.auth.PhoneAuthProvider.PROVIDER_ID,
				//firebaseui.auth.AnonymousAuthProvider.PROVIDER_ID,
			],
			// tosUrl and privacyPolicyUrl accept either url string or a callback
			// function.
			// Terms of service url/callback.
			//tosUrl: '<your-tos-url>',
			// Privacy policy url/callback.
			//privacyPolicyUrl: function () {
			//window.location.assign('<your-privacy-policy-url>');
			//},
		};

		//eslint-disable-next-line @typescript-eslint/no-explicit-any
		async function authUser(user: any) {
			try {
				const idToken = await user.getIdToken(true),
					wellreadAuthPayload = {
						idToken,
					};
				await fetch('/api/auth/callback', {
					method: 'POST',
					body: JSON.stringify(wellreadAuthPayload),
					headers: {
						'Content-Type': 'application/json',
					},
				});
				location.href = '/';
			} catch (e) {
				console.error(e);
			}
		}

		// Initialize the FirebaseUI Widget using Firebase.
		//eslint-disable-next-line @typescript-eslint/no-explicit-any
		const ui = new (window as any).firebaseui.auth.AuthUI(firebase.auth());
		// The start method will wait until the DOM is loaded.
		ui.start('#firebaseui-auth-container', uiConfig);
	});
</script>
