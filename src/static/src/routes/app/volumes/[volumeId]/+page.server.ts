import { makeApiUrl } from '$lib/proxy';
import type { ReadingHistory, Volume } from '$lib/types';
import { error } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch, params }) => {
	const json = (url: string) => fetch(makeApiUrl(url)).then(async (res) => ({ body: await res.json(), ok: res.ok })),
		volume = await json(`/api/volumes/${params.volumeId}`);

	// if they tried querying for an invalid volume
	if (!volume.ok) {
		throw error(404, 'Volume not found.');
	}

	return {
		volume: volume.body,
		readingHistory: (await json(`/api/volumes/${params.volumeId}/history`)).body,
	} as {
		volume: Volume;
		readingHistory: ReadingHistory[];
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	deleteReadingHistory: async ({ request, fetch }) => {
		const data = await request.formData();

		await fetch(makeApiUrl(`/api/history/${data.get('readingHistoryId')}`), {
			method: 'DELETE',
		});
	},
};
