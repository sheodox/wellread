import { makeApiUrl } from '$lib/proxy';
import type { Series } from '$lib/types';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch }) => {
	return {
		series: await (await fetch(makeApiUrl(`/api/series`))).json(),
	} as {
		series: Series[];
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	newSeries: async ({ request, fetch }) => {
		const name = (await request.formData()).get('name');
		const series: Series = await fetch(makeApiUrl(`/api/series`), {
			method: 'POST',
			body: JSON.stringify({
				name,
			}),
			headers: {
				'Content-Type': 'application/json',
			},
		}).then((res) => res.json());

		throw redirect(303, `/app/series/${series.id}`);
	},
};
