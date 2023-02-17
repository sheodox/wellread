import { makeApiUrl } from '$lib/proxy';
import type { PagedResponse, Volume } from '$lib/types';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch }) => {
	const loggedIn = (await fetch(makeApiUrl('/api/auth/status'))).ok;

	return {
		loggedIn,
		reading: await (await fetch(makeApiUrl('/api/volumes?status=reading'))).json(),
	} as {
		loggedIn: boolean;
		reading: PagedResponse<Volume[]>;
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	logout: async (event) => {
		event.cookies.delete('session');
		throw redirect(303, '/');
	},
};
