import { makeApiUrl } from '$lib/proxy';
import type { Series } from '$lib/types';
import { collectFormValues } from '$lib/util';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch, params }) => {
	const json = (url: string) => fetch(makeApiUrl(url)).then((res) => res.json());

	return {
		series: await json(`/api/series/${params.seriesId}`),
	} as {
		series: Series;
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	update: async ({ request, fetch, params }) => {
		const data = collectFormValues(await request.formData());

		await fetch(makeApiUrl(`/api/series/${params.seriesId}`), {
			method: 'PATCH',
			body: JSON.stringify(data),
			headers: {
				'Content-Type': 'application/json',
			},
		});

		throw redirect(303, `/app/series/${params.seriesId}`);
	},
};
