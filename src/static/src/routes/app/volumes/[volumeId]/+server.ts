import { makeApiUrl } from '$lib/proxy';
import type { RequestHandler } from './$types';

export const DELETE = (async ({ params, fetch }) => {
	return await fetch(makeApiUrl(`/api/volumes/${params.volumeId}`), {
		method: 'DELETE',
	});
}) satisfies RequestHandler;
