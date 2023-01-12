import { makeApiUrl } from '$lib/proxy';
import type { Volume } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch }) => {
	return {
		volumes: await (await fetch(makeApiUrl(`/api/volumes`))).json(),
	} as {
		volumes: Volume[];
	};
}) satisfies PageServerLoad;
