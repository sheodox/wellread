import { makeApiUrl } from '$lib/proxy';
import { redirect, type Handle, type HandleFetch } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	event.locals.isLoggedIn = (
		await fetch(makeApiUrl('/api/auth/status'), {
			headers: {
				cookie: event.request.headers.get('cookie') ?? '',
			},
		})
	).ok;

	// if a non-logged in user goes to the app, redirect to the home page so they can log in
	if (!event.locals.isLoggedIn && event.route.id?.startsWith('/app')) {
		throw redirect(303, '/');
	}

	// redirect users to the app if they go to the home page
	if (event.locals.isLoggedIn && event.url.pathname === '/') {
		throw redirect(303, '/app');
	}

	const response = await resolve(event);
	return response;
}) satisfies Handle;

export const handleFetch = (({ event, request, fetch }) => {
	if (request.url.startsWith('http://api:5004')) {
		// pass cookies along to requests to the API within docker
		request.headers.set('cookie', event.request.headers.get('cookie') ?? '');
	}
	return fetch(request);
}) satisfies HandleFetch;
