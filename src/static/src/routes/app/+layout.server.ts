import { makeApiUrl } from '$lib/proxy';
import type { LayoutServerLoad } from './$types';

export const load = (async ({ fetch }) => {
	const loggedIn = (await fetch(makeApiUrl('/api/auth/status'))).ok;

	return {
		loggedIn: loggedIn as boolean,
	};
}) satisfies LayoutServerLoad;
