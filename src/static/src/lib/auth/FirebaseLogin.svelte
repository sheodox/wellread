<div id="firebaseui-auth-container" />

<script lang="ts">
	import firebase from 'firebase/compat/app';
	import * as firebaseui from 'firebaseui';
	import 'firebaseui/dist/firebaseui.css';

	fetch('/auth/firebase-config').then(async (res) => {
		const { body: firebaseConfig } = await res.json();
		firebase.initializeApp(firebaseConfig);

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

		async function authUser(user: firebase.User) {
			try {
				const idToken = await user.getIdToken(true),
					wellreadAuthPayload = {
						idToken,
					};
				await fetch('/auth/callback', {
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
		const ui = new firebaseui.auth.AuthUI(firebase.auth());
		// The start method will wait until the DOM is loaded.
		ui.start('#firebaseui-auth-container', uiConfig);
	});
</script>
