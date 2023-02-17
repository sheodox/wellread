import { makeApiUrl } from '$lib/proxy';
import { error, redirect } from '@sveltejs/kit';
import type { PagedResponse, Series, Volume } from '$lib/types';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch, params, url }) => {
	const series = await fetch(makeApiUrl(`/api/series/${params.seriesId}`));

	if (!series.ok) {
		throw error(404, 'Series not found');
	}

	const pageNum = url.searchParams.get('page') ?? '1';

	return {
		volumes: (await fetch(makeApiUrl(`/api/volumes?seriesId=${params.seriesId}&page=${pageNum}`))).json(),
		series: series.json(),
	} as {
		volumes: Promise<PagedResponse<Volume[]>>;
		series: Promise<Series>;
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	newVolume: async ({ request, fetch, params }) => {
		const name = (await request.formData()).get('name');
		const volume: Volume = await fetch(makeApiUrl(`/api/series/${params.seriesId}/volumes`), {
			method: 'POST',
			body: JSON.stringify({
				name,
			}),
			headers: {
				'Content-Type': 'application/json',
			},
		}).then((res) => res.json());

		throw redirect(303, `/app/volumes/${volume.id}`);
	},
};
