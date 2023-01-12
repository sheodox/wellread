import { makeApiUrl } from '$lib/proxy';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch }) => {
	return {
		firebaseConfig: (await (await fetch(makeApiUrl('/api/auth/firebase-config'))).json()) as Record<string, string>,
	};
}) satisfies PageServerLoad;
